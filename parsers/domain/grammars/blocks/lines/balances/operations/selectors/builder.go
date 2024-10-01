package selectors

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/operations/selectors/chains"
)

type builder struct {
	chain chains.Chain
	isNot bool
}

func createBuilder() Builder {
	out := builder{
		chain: nil,
		isNot: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithChain adds a chain to the builder
func (app *builder) WithChain(chain chains.Chain) Builder {
	app.chain = chain
	return app
}

// IsNot flags the builder as isNot
func (app *builder) IsNot() Builder {
	app.isNot = true
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Selector instance")
	}

	return createSelector(
		app.chain,
		app.isNot,
	), nil
}
