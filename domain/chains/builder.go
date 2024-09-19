package chains

import (
	"errors"

	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     nfts.NFT
	action      Action
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		grammar:     nil,
		action:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar nfts.NFT) Builder {
	app.grammar = grammar
	return app
}

// WithAction adds an action to the builder
func (app *builder) WithAction(action Action) Builder {
	app.action = action
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Chain instance")
	}

	if app.action == nil {
		return nil, errors.New("the action is mandatory in order to build a Chain instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.grammar.Hash().Bytes(),
		app.action.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createChain(*pHash, app.grammar, app.action), nil
}
