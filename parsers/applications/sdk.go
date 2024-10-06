package applications

import (
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/queries"
	"github.com/steve-care-software/steve/parsers/domain/walkers/languages"
)

// NewApplication creates a new application
func NewApplication() Application {
	elementsAdapter := instructions.NewElementsAdapter()
	astAdapter := asts.NewAdapter()
	queryAdapter, _ := queries.NewAdapterFactory().Create()
	tokensBuilder := instructions.NewTokensBuilder()
	return createApplication(
		elementsAdapter,
		astAdapter,
		queryAdapter,
		tokensBuilder,
	)
}

// Application represents the interpreter application
type Application interface {
	// Execute executes the parser
	Execute(input []byte, grammar grammars.Grammar, ins languages.Element) (any, []byte, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
