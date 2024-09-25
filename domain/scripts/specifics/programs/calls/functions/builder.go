package functions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions/parameters"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	parameters  parameters.Parameters
	isEngine    bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		parameters:  nil,
		isEngine:    false,
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

// WithParameters add parameters to the builder
func (app *builder) WithParameters(parameters parameters.Parameters) Builder {
	app.parameters = parameters
	return app
}

// IsEngine flags the builder as an engine
func (app *builder) IsEngine() Builder {
	app.isEngine = true
	return app
}

// Now builds a new Function instance
func (app *builder) Now() (Function, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Function instance")
	}

	if app.parameters == nil {
		return nil, errors.New("the parameters is mandatory in order to build a Function instance")
	}

	isEngine := "false"
	if app.isEngine {
		isEngine = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.parameters.Hash().Bytes(),
		[]byte(isEngine),
	})

	if err != nil {
		return nil, err
	}

	return createFunction(
		*pHash,
		app.name,
		app.parameters,
		app.isEngine,
	), nil
}
