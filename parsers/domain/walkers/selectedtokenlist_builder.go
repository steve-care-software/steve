package walkers

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

type selectedTokenListBuilder struct {
	name  string
	chain chains.Chain
	node  Node
}

func createSelectedTokenListBuilder() SelectedTokenListBuilder {
	out := selectedTokenListBuilder{
		name:  "",
		chain: nil,
		node:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *selectedTokenListBuilder) Create() SelectedTokenListBuilder {
	return createSelectedTokenListBuilder()
}

// WithName adds a name to the builder
func (app *selectedTokenListBuilder) WithName(name string) SelectedTokenListBuilder {
	app.name = name
	return app
}

// WithChain adds a chain to the builder
func (app *selectedTokenListBuilder) WithChain(chain chains.Chain) SelectedTokenListBuilder {
	app.chain = chain
	return app
}

// WithNode adds a node to the builder
func (app *selectedTokenListBuilder) WithNode(node Node) SelectedTokenListBuilder {
	app.node = node
	return app
}

// Now builds a new SelectedTokenList instance
func (app *selectedTokenListBuilder) Now() (SelectedTokenList, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a SelectedTokenList instance")
	}

	if app.chain != nil && app.node != nil {
		return createSelectedTokenListWithChainAndNode(
			app.name,
			app.chain,
			app.node,
		), nil
	}

	if app.chain != nil {
		return createSelectedTokenListWithChain(
			app.name,
			app.chain,
		), nil
	}

	if app.node != nil {
		return createSelectedTokenListWithNode(
			app.name,
			app.node,
		), nil
	}

	return createSelectedTokenList(app.name), nil
}
