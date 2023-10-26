package executions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

type builder struct {
	layer  layers.LayerInput
	values layers.ValueAssignments
}

func createBuilder() Builder {
	out := builder{
		layer:  nil,
		values: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer layers.LayerInput) Builder {
	app.layer = layer
	return app
}

// WithValues add values to the builder
func (app *builder) WithValues(values layers.ValueAssignments) Builder {
	app.values = values
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build an Execution instance")
	}

	if app.values != nil {
		return createExecutionWithValues(app.layer, app.values), nil
	}

	return createExecution(app.layer), nil
}
