package asts

import (
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/grammars/rules"
)

// NewParserAdapter creates a new parser adapter
func NewParserAdapter() ParserAdapter {
	grammarAdapter := grammars.NewParserAdapter()
	builder := NewBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	tokensBuilder := instructions.NewTokensBuilder()
	tokenBuilder := instructions.NewTokenBuilder()
	elementsBuilder := instructions.NewElementsBuilder()
	elementBuilder := instructions.NewElementBuilder()
	ruleBuilder := rules.NewRuleBuilder()
	syscallBuilder := instructions.NewSyscallBuilder()
	parametersBuilder := instructions.NewParametersBuilder()
	parameterBuilder := instructions.NewParameterBuilder()
	valueBuilder := instructions.NewValueBuilder()
	referenceBuilder := instructions.NewReferenceBuilder()
	return createParserAdapter(
		grammarAdapter,
		builder,
		instructionsBuilder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		ruleBuilder,
		syscallBuilder,
		parametersBuilder,
		parameterBuilder,
		valueBuilder,
		referenceBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// ParserAdapter represents the ast parser adapter
type ParserAdapter interface {
	// ToAST takes the grammar and input and converts them to a ast instance and the remaining data
	ToAST(grammar grammars.Grammar, input []byte) (AST, []byte, error)

	// ToASTWithRoot creates a ast but changes the root block of the grammar
	ToASTWithRoot(grammar grammars.Grammar, rootBlockName string, input []byte) (AST, []byte, error)
}

// Builder represents the ast builder
type Builder interface {
	Create() Builder
	WithRoot(root instructions.Element) Builder
	Now() (AST, error)
}

// AST represents a ast
type AST interface {
	Root() instructions.Element
}
