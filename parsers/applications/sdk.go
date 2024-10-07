package applications

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/walkers/elements"
)

// NewBuilder creates a new application builder
func NewBuilder() Builder {
	elementsAdapter := instructions.NewElementsAdapter()
	astAdapter := asts.NewAdapter()
	elementAdapter := elements.NewAdapter()
	tokensBuilder := instructions.NewTokensBuilder()
	return createBuilder(
		elementsAdapter,
		astAdapter,
		elementAdapter,
		tokensBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithElement(ins elements.Element) Builder
	Now() (Application, error)
}

// Application represents the interpreter application
type Application interface {
	// Execute executes the parser
	Execute(input []byte, grammar grammars.Grammar) (any, []byte, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
