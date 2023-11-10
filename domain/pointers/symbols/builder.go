package symbols

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers/symbols/kinds"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	kind        kinds.Kind
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

// Now builds a new Symbol instance
func (app *builder) Now() (Symbol, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Symbol instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Symbol instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSymbol(*pHash, app.name, app.kind), nil
}
