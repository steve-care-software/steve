package selectors

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type selectorBuilder struct {
	chain chains.Chain
	isNot bool
}

func createSelectorBuilder() SelectorBuilder {
	out := selectorBuilder{
		chain: nil,
		isNot: false,
	}

	return &out
}

// Create initializes the selectorBuilder
func (app *selectorBuilder) Create() SelectorBuilder {
	return createSelectorBuilder()
}

// WithChain adds a chain to the selectorBuilder
func (app *selectorBuilder) WithChain(chain chains.Chain) SelectorBuilder {
	app.chain = chain
	return app
}

// IsNot flags the selectorBuilder as isNot
func (app *selectorBuilder) IsNot() SelectorBuilder {
	app.isNot = true
	return app
}

// Now builds a new Selector instance
func (app *selectorBuilder) Now() (Selector, error) {
	if app.chain == nil {
		return nil, errors.New("the chain is mandatory in order to build a Selector instance")
	}

	return createSelector(
		app.chain,
		app.isNot,
	), nil
}
