package asts

import (
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	grammarAdapter := grammars.NewAdapter()
	builder := NewBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	tokensBuilder := instructions.NewTokensBuilder()
	tokenBuilder := instructions.NewTokenBuilder()
	elementsBuilder := instructions.NewElementsBuilder()
	elementBuilder := instructions.NewElementBuilder()
	constantBuilder := instructions.NewConstantBuilder()
	return createAdapter(
		grammarAdapter,
		builder,
		instructionsBuilder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		constantBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the adapter
type Adapter interface {
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
