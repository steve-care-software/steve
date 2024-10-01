package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type application struct {
	elementsAdapter instructions.ElementsAdapter
	astAdapter      asts.Adapter
}

func createApplication(
	elementsAdapter instructions.ElementsAdapter,
	astAdapter asts.Adapter,
) Application {
	out := application{
		elementsAdapter: elementsAdapter,
		astAdapter:      astAdapter,
	}

	return &out
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
	if element.IsRule() {
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
