package parameters

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters/kinds"
)

type parameterBuilder struct {
	hashAdapter hash.Adapter
	name        string
	kind        kinds.Kind
}

func createParameterBuilder(
	hashAdapter hash.Adapter,
) ParameterBuilder {
	out := parameterBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// WithKind adds a kind to the builder
func (app *parameterBuilder) WithKind(kind kinds.Kind) ParameterBuilder {
	app.kind = kind
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Parameter instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createParameter(*pHash, app.name, app.kind), nil
}
