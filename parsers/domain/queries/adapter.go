package queries

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type adapter struct {
	astAdapter            asts.Adapter
	builder               Builder
	grammarElementBuilder elements.ElementBuilder
	chainBuilder          chains.Builder
	tokenBuilder          chains.TokenBuilder
	elementBuilder        chains.ElementBuilder
	grammar               grammars.Grammar
}

func createAdapter(
	astAdapter asts.Adapter,
	builder Builder,
	grammarElementBuilder elements.ElementBuilder,
	chainBuilder chains.Builder,
	tokenBuilder chains.TokenBuilder,
	elementBuilder chains.ElementBuilder,
	grammar grammars.Grammar,
) Adapter {
	out := adapter{
		astAdapter:            astAdapter,
		builder:               builder,
		grammarElementBuilder: grammarElementBuilder,
		chainBuilder:          chainBuilder,
		tokenBuilder:          tokenBuilder,
		elementBuilder:        elementBuilder,
		grammar:               grammar,
	}

	return &out
}

// ToQuery converts bytes to a Query instance
func (app *adapter) ToQuery(input []byte) (Query, []byte, error) {
	ast, retRemaining, err := app.astAdapter.ToAST(app.grammar, input)
	if err != nil {
		return nil, nil, err
	}

	root := ast.Root()
	retQuery, err := app.query(root)
	if err != nil {
		return nil, nil, err
	}

	return retQuery, retRemaining, nil
}

func (app *adapter) query(element instructions.Element) (Query, error) {
	tokens, err := app.elementToTokens(element)
	if err != nil {
		return nil, err
	}

	retHeadToken, err := tokens.Fetch("head", 0)
	if err != nil {
		return nil, err
	}

	version, name, err := app.head(retHeadToken)
	if err != nil {
		return nil, err
	}

	retChainLineToken, err := tokens.Fetch("chainLine", 0)
	if err != nil {
		return nil, err
	}

	chain, err := app.chainLine(retChainLineToken)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithVersion(version).
		WithName(name).
		WithChain(chain).
		Now()
}

func (app *adapter) chainLine(token instructions.Token) (chains.Chain, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	retChainToken, err := tokens.Fetch("chain", 0)
	if err != nil {
		return nil, err
	}

	return app.chain(retChainToken)
}

func (app *adapter) chain(token instructions.Token) (chains.Chain, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	retNameToken, err := tokens.Fetch("elementName", 0)
	if err != nil {
		return nil, err
	}

	retGrammarElement, err := app.grammarElement(retNameToken)
	if err != nil {
		return nil, err
	}

	builder := app.chainBuilder.Create().WithElement(retGrammarElement)
	retTokenToken, err := tokens.Fetch("token", 0)
	if err == nil {
		retToken, err := app.token(retTokenToken)
		if err != nil {
			return nil, err
		}

		builder.WithToken(retToken)
	}

	return builder.Now()
}

func (app *adapter) token(token instructions.Token) (chains.Token, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	retIndexToken, err := tokens.Fetch("index", 0)
	if err != nil {
		return nil, err
	}

	index, err := app.numbersInstruction(retIndexToken)
	if err != nil {
		return nil, err
	}

	builder := app.tokenBuilder.Create().WithIndex(uint(index))
	retElementToken, err := tokens.Fetch("element", 0)
	if err == nil {
		retElement, err := app.element(retElementToken)
		if err != nil {
			return nil, err
		}

		builder.WithElement(retElement)
	}

	return builder.Now()
}

func (app *adapter) element(token instructions.Token) (chains.Element, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	retIndexToken, err := tokens.Fetch("index", 0)
	if err != nil {
		return nil, err
	}

	index, err := app.numbersInstruction(retIndexToken)
	if err != nil {
		return nil, err
	}

	builder := app.elementBuilder.Create().WithIndex(uint(index))
	retArrowChainToken, err := tokens.Fetch("arrowChain", 0)
	if err == nil {
		retArrowChainTokens, err := app.tokenToFirstInstructionTokens(retArrowChainToken)
		if err != nil {
			return nil, err
		}

		retArrowChainToken, err := retArrowChainTokens.Fetch("chain", 0)
		if err != nil {
			return nil, err
		}

		retChain, err := app.chain(retArrowChainToken)
		if err != nil {
			return nil, err
		}

		builder.WithChain(retChain)
	}

	return builder.Now()
}

func (app *adapter) grammarElement(token instructions.Token) (elements.Element, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return nil, err
	}

	builder := app.grammarElementBuilder.Create()
	retBlockToken, err := tokens.Fetch("variableName", 0)
	if err == nil {
		builder.WithBlock(string(retBlockToken.Value()))
	}

	retConstantToken, err := tokens.Fetch("constantName", 0)
	if err == nil {
		builder.WithConstant(string(retConstantToken.Value()))
	}

	retRuleToken, err := tokens.Fetch("ruleName", 0)
	if err == nil {
		builder.WithRule(string(retRuleToken.Value()))
	}

	return builder.Now()
}

func (app *adapter) head(token instructions.Token) (uint, string, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return 0, "", err
	}

	retVersionToken, err := tokens.Fetch("versionInstruction", 0)
	if err != nil {
		return 0, "", err
	}

	versionNumber, err := app.numbersInstruction(retVersionToken)
	if err != nil {
		return 0, "", err
	}

	retNameToken, err := tokens.Fetch("nameInstruction", 0)
	if err != nil {
		return 0, "", err
	}

	nameStr, err := app.nameInstruction(retNameToken)
	if err != nil {
		return 0, "", err
	}

	return uint(versionNumber), nameStr, nil
}

func (app *adapter) nameInstruction(token instructions.Token) (string, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return "", err
	}

	retNameToken, err := tokens.Fetch("variableName", 0)
	if err != nil {
		return "", err
	}

	return string(retNameToken.Value()), nil
}

func (app *adapter) numbersInstruction(token instructions.Token) (int, error) {
	tokens, err := app.tokenToFirstInstructionTokens(token)
	if err != nil {
		return 0, err
	}

	retNumberToken, err := tokens.Fetch("numbers", 0)
	if err != nil {
		return 0, err
	}

	retBytes := retNumberToken.Value()
	return strconv.Atoi(string(retBytes))
}

func (app *adapter) tokenToFirstInstructionTokens(token instructions.Token) (instructions.Tokens, error) {
	retElement, err := token.Elements().Fetch(0)
	if err != nil {
		return nil, err
	}

	return app.elementToTokens(retElement)
}

func (app *adapter) elementToTokens(element instructions.Element) (instructions.Tokens, error) {
	if element.IsConstant() {
		str := fmt.Sprintf("the element (name: %s) was expected to contain an Instruction", element.Name())
		return nil, errors.New(str)
	}

	return element.Instruction().Tokens(), nil
}
