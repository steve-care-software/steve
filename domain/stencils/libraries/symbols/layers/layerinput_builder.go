package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type layerInputBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	layer       Layer
}

func createLayerInputBuilder(
	hashAdapter hash.Adapter,
) LayerInputBuilder {
	out := layerInputBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		layer:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerInputBuilder) Create() LayerInputBuilder {
	return createLayerInputBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *layerInputBuilder) WithVariable(variable string) LayerInputBuilder {
	app.variable = variable
	return app
}

// WithLayer adds a layer to the builder
func (app *layerInputBuilder) WithLayer(layer Layer) LayerInputBuilder {
	app.layer = layer
	return app
}

// Now builds a new LayerInput instance
func (app *layerInputBuilder) Now() (LayerInput, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a LayerInput instance")
	}

	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a LayerInput instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.layer.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLayerInput(*pHash, app.variable, app.layer), nil
}
