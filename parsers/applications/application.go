package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
	"github.com/steve-care-software/steve/parsers/domain/queries"
)

type application struct {
	elementsAdapter instructions.ElementsAdapter
	astAdapter      asts.Adapter
	queryAdapter    queries.Adapter
	tokensBuilder   instructions.TokensBuilder
}

func createApplication(
	elementsAdapter instructions.ElementsAdapter,
	astAdapter asts.Adapter,
	queryAdapter queries.Adapter,
	tokensBuilder instructions.TokensBuilder,
) Application {
	out := application{
		elementsAdapter: elementsAdapter,
		astAdapter:      astAdapter,
		queryAdapter:    queryAdapter,
		tokensBuilder:   tokensBuilder,
	}

	return &out
}

// Execute executes the parser application
func (app *application) Execute(input []byte, grammar grammars.Grammar, ins Element) (any, []byte, error) {
	ast, retRemaining, err := app.astAdapter.ToAST(grammar, input)
	if err != nil {
		return nil, nil, err
	}

	root := ast.Root()
	retIns, err := app.element(root, ins)
	if err != nil {
		return nil, nil, err
	}

	return retIns, retRemaining, nil
}

func (app *application) element(element instructions.Element, ins Element) (any, error) {
	if element.IsConstant() {
		value := element.Constant().Value()
		return app.callElementFn(value, ins.ElementFn)
	}

	elementName := element.Name()
	tokens := element.Instruction().Tokens()
	if ins.TokenList != nil {
		ptrTokenList := ins.TokenList
		output, err := app.tokenList(elementName, tokens.List(), *ptrTokenList)
		if err != nil {
			return nil, err
		}

		return app.callElementFn(output, ins.ElementFn)
	}

	value := tokens.Value()
	return app.callElementFn(value, ins.ElementFn)
}

func (app *application) callElementFn(value any, fn ElementFn) (any, error) {
	if fn == nil {
		return value, nil
	}

	return fn(value)
}

func (app *application) tokenList(
	elementName string,
	tokensList []instructions.Token,
	ins TokenList,
) (any, error) {
	output := map[string][]any{}
	for _, oneToken := range tokensList {
		name := oneToken.Name()
		if chosenTokenList, ok := ins.List[name]; ok {
			retValue, err := app.chosenTokenList(name, tokensList, chosenTokenList)
			if err != nil {
				return nil, err
			}

			if _, ok := output[name]; !ok {
				output[name] = []any{}
			}

			output[name] = append(output[name], retValue)
		}

	}

	return ins.MapFn(elementName, output)
}

func (app *application) token(token instructions.Token, ins Token) (any, error) {
	output := []any{}
	elementsList := token.Elements().List()
	for _, oneElement := range elementsList {
		if ins.Next == nil {
			output = append(output, oneElement.Value())
			continue
		}

		retValue, err := app.element(oneElement, *ins.Next)
		if err != nil {
			return nil, err
		}

		output = append(output, retValue)
	}

	return ins.ListFn(output)
}

func (app *application) chosenTokenList(
	elementName string,
	tokensList []instructions.Token,
	ins ChosenTokenList,
) (any, error) {
	tokensIns, err := app.tokensBuilder.Create().
		WithList(tokensList).
		Now()

	if err != nil {
		return nil, err
	}

	var tokenIns instructions.Token
	var elementIns instructions.Element
	if ins.SelectorScript != nil {
		retTokensList, retToken, retElement, err := app.query(
			tokensIns,
			ins.SelectorScript,
		)

		if err != nil {
			return nil, err
		}

		tokensIns = nil
		if len(retTokensList) > 0 {
			retTokens, err := app.tokensBuilder.Create().
				WithList(retTokensList).
				Now()

			if err != nil {
				return nil, err
			}

			tokensIns = retTokens
		}

		if retToken != nil {
			tokenIns = retToken
		}

		if retElement != nil {
			elementIns = retElement
		}
	}

	if ins.Node == nil {
		if tokensIns != nil {
			return tokensIns.Value(), nil
		}

		if tokenIns != nil {
			return tokenIns.Value(), nil
		}

		return elementIns.Value(), nil
	}

	return app.node(
		elementName,
		tokensIns,
		tokenIns,
		elementIns,
		*ins.Node,
	)
}

func (app *application) node(
	elementName string,
	tokens instructions.Tokens,
	token instructions.Token,
	element instructions.Element,
	ins Node,
) (any, error) {
	if tokens != nil {
		if ins.TokenList == nil {
			str := fmt.Sprintf("the element (%s) was expected to contain a Tokens instance after its selection, but the TokenList parser did not exists", elementName)
			return nil, errors.New(str)
		}

		return app.tokenList(elementName, tokens.List(), *ins.TokenList)
	}

	if token != nil {
		if ins.Token == nil {
			str := fmt.Sprintf("the element (%s) was expected to contain a Token instance after its selection, but the Token parser did not exists", elementName)
			return nil, errors.New(str)
		}

		return app.token(token, *ins.Token)
	}

	if ins.Element == nil {
		str := fmt.Sprintf("the element (%s) was expected to contain an Element instance after its selection, but the Element parser did not exists", elementName)
		return nil, errors.New(str)
	}

	return app.element(element, *ins.Element)
}

func (app *application) query(tokens instructions.Tokens, script []byte) ([]instructions.Token, instructions.Token, instructions.Element, error) {
	query, remaining, err := app.queryAdapter.ToQuery(script)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(remaining) > 0 {
		str := fmt.Sprintf("the script (%s) contains a remaining (%s)", script, remaining)
		return nil, nil, nil, errors.New(str)
	}

	chain := query.Chain()
	return tokens.Select(chain)
}

// Suites executes all the test suites of the grammar
func (app *application) Suites(grammar grammars.Grammar) error {
	blocksList := grammar.Blocks().List()
	for _, oneBlock := range blocksList {
		if !oneBlock.HasSuites() {
			continue
		}

		blockName := oneBlock.Name()
		suitesList := oneBlock.Suites().List()
		for idx, oneSuite := range suitesList {
			err := app.interpretSuite(
				grammar,
				blockName,
				oneSuite,
			)

			prefix := fmt.Sprintf("block (name: %s) index (%d) suite (%s)", blockName, idx, oneSuite.Name())
			if oneSuite.IsFail() {
				if err == nil {
					str := fmt.Sprintf("%s: the suite was expected to FAIL but succeeded!", prefix)
					return errors.New(str)
				}

				continue
			}

			if err != nil {
				str := fmt.Sprintf("%s the suite was expected to SUCCEED but failed --- error: %s", prefix, err.Error())
				return errors.New(str)
			}
		}
	}
	return nil
}

func (app *application) interpretSuite(
	grammar grammars.Grammar,
	blockName string,
	suite suites.Suite,
) error {
	ast, retRemaining, err := app.astAdapter.ToASTWithRoot(
		grammar,
		blockName,
		suite.Input(),
	)

	if err != nil {
		return err
	}

	if len(retRemaining) != 0 {
		str := fmt.Sprintf("the bytes (%s) were remaining", retRemaining)
		return errors.New(str)
	}

	return app.execute(ast)
}

func (app *application) interpretInstruction(
	instruction instructions.Instruction,
) error {
	tokens := instruction.Tokens()
	return app.interpretTokens(
		tokens,
	)
}

func (app *application) interpretTokens(
	tokens instructions.Tokens,
) error {
	list := tokens.List()
	for _, oneToken := range list {
		err := app.interpretToken(
			tokens,
			oneToken,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretToken(
	currentTokens instructions.Tokens,
	token instructions.Token,
) error {
	elements := token.Elements()
	return app.interpretElements(
		currentTokens,
		elements,
	)
}

func (app *application) interpretElements(
	currentTokens instructions.Tokens,
	elements instructions.Elements,
) error {
	list := elements.List()
	for _, oneElement := range list {
		err := app.interpretElement(
			currentTokens,
			oneElement,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) interpretElement(
	currentTokens instructions.Tokens,
	element instructions.Element,
) error {
	if element.IsConstant() {
		return nil
	}

	instruction := element.Instruction()
	return app.interpretInstruction(
		instruction,
	)
}

func (app *application) execute(ast asts.AST) error {
	root := ast.Root()
	return app.interpretElement(
		nil,
		root,
	)
}
