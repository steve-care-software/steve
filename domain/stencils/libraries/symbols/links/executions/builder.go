package executions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

type builder struct {
	hashAdapter hash.Adapter
	layer       layers.LayerInput
	values      layers.ValueAssignments
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layer:       nil,
		values:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{
		app.layer.Hash().Bytes(),
	}

	if app.values != nil {
		data = append(data, app.values.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.values != nil {
		return createExecutionWithValues(*pHash, app.layer, app.values), nil
	}

	return createExecution(*pHash, app.layer), nil
}
