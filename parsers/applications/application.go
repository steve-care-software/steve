package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
	"github.com/steve-care-software/steve/parsers/domain/walkers"
)

type application struct {
	elementsAdapter instructions.ElementsAdapter
	astAdapter      asts.Adapter
	tokensBuilder   instructions.TokensBuilder
	walker          walkers.Walker
}

func createApplication(
	elementsAdapter instructions.ElementsAdapter,
	astAdapter asts.Adapter,
	tokensBuilder instructions.TokensBuilder,
	walker walkers.Walker,
) Application {
	out := application{
		elementsAdapter: elementsAdapter,
		astAdapter:      astAdapter,
		tokensBuilder:   tokensBuilder,
		walker:          walker,
	}

	return &out
}

// Execute executes the parser application
func (app *application) Execute(input []byte, grammar grammars.Grammar) (any, []byte, error) {
	if app.walker == nil {
		return nil, nil, errors.New("the application cannot Execute because it doesn't contain a Walker instance")
	}
	ast, retRemaining, err := app.astAdapter.ToAST(grammar, input)
	if err != nil {
		return nil, nil, err
	}

	root := ast.Root()
	retIns, err := app.element(root, app.walker)
	if err != nil {
		return nil, nil, err
	}

	return retIns, retRemaining, nil
}

func (app *application) element(element instructions.Element, ins walkers.Walker) (any, error) {
	if element.IsConstant() {
		value := element.Constant().Value()
		return app.callElementFn(value, ins.Fn())
	}

	elementName := element.Name()
	tokens := element.Instruction().Tokens()
	if ins.HasList() {
		output, err := app.tokenList(elementName, tokens.List(), ins.List())
		if err != nil {
			return nil, err
		}

		return app.callElementFn(output, ins.Fn())
	}

	value := tokens.Value()
	return app.callElementFn(value, ins.Fn())
}

func (app *application) callElementFn(value any, fn walkers.ElementFn) (any, error) {
	if fn == nil {
		return value, nil
	}

	return fn(value)
}

func (app *application) tokenList(
	elementName string,
	tokensList []instructions.Token,
	ins walkers.TokenList,
) (any, error) {
	output := map[string][]any{}
	for _, oneToken := range tokensList {
		name := oneToken.Name()
		selectedTokenList, err := ins.Fetch(name)
		if err != nil {
			continue
		}

		if name == "pointSuite" {
			fmt.Printf("\n ---+ %s, %s, %v, %v\n", elementName, name, tokensList, ins)
		}

		retValue, err := app.selectedTokenList(name, tokensList, selectedTokenList)
		if err != nil {
			return nil, err
		}

		if _, ok := output[name]; !ok {
			output[name] = []any{}
		}

		output[name] = append(output[name], retValue)
	}

	return ins.Fn()(elementName, output)
}

func (app *application) token(token instructions.Token, ins walkers.Token) (any, error) {
	output := []any{}
	elementsList := token.Elements().List()
	for _, oneElement := range elementsList {
		if !ins.HasNext() {
			output = append(output, oneElement.Value())
			continue
		}

		retValue, err := app.element(oneElement, ins.Next())
		if err != nil {
			return nil, err
		}

		output = append(output, retValue)
	}

	return ins.Fn()(output)
}

func (app *application) selectedTokenList(
	elementName string,
	tokensList []instructions.Token,
	ins walkers.SelectedTokenList,
) (any, error) {
	tokensIns, err := app.tokensBuilder.Create().
		WithList(tokensList).
		Now()

	if err != nil {
		return nil, err
	}

	var tokenIns instructions.Token
	var elementIns instructions.Element
	if ins.HasChain() {
		chain := ins.Chain()
		retTokensList, retToken, retElement, err := tokensIns.Select(chain)
		if err != nil {
			str := fmt.Sprintf("the element (name: %s) contains a script that does not match the provided Token instance", elementName)
			return nil, errors.New(str)
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

	if !ins.HasNode() {
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
		ins.Node(),
	)
}

func (app *application) node(
	elementName string,
	tokens instructions.Tokens,
	token instructions.Token,
	element instructions.Element,
	ins walkers.Node,
) (any, error) {
	if tokens != nil {
		if !ins.IsTokenList() {
			str := fmt.Sprintf("the element (%s) was expected to contain a Tokens instance after its selection, but the TokenList parser did not exists", elementName)
			return nil, errors.New(str)
		}

		return app.tokenList(elementName, tokens.List(), ins.TokenList())
	}

	if token != nil {
		if !ins.IsToken() {
			str := fmt.Sprintf("the element (%s) was expected to contain a Token instance after its selection, but the Token parser did not exists", elementName)
			return nil, errors.New(str)
		}

		return app.token(token, ins.Token())
	}

	if !ins.IsElement() {
		str := fmt.Sprintf("the element (%s) was expected to contain an Element instance after its selection, but the Element parser did not exists", elementName)
		return nil, errors.New(str)
	}

	return app.element(element, ins.Element())
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
