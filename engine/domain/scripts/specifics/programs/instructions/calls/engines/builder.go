package engines

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/functions"
)

type builder struct {
	hashAdapter hash.Adapter
	pScope      *uint8
	function    functions.Function
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pScope:      nil,
		function:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithScope adds a scope to the builder
func (app *builder) WithScope(scope uint8) Builder {
	app.pScope = &scope
	return app
}

// WithFunction adds a function to the builder
func (app *builder) WithFunction(function functions.Function) Builder {
	app.function = function
	return app
}

// Now builds a new Engine instance
func (app *builder) Now() (Engine, error) {
	if app.pScope == nil {
		return nil, errors.New("the scope is mandatory in order to build an Engine instance")
	}

	if app.function == nil {
		return nil, errors.New("the function is mandatory in order to build an Engine instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		{*app.pScope},
		app.function.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createEngine(
		*pHash,
		*app.pScope,
		app.function,
	), nil
}
