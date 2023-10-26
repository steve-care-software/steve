package layers

import "errors"

type layerInputBuilder struct {
	variable string
	layer    Layer
}

func createLayerInputBuilder() LayerInputBuilder {
	out := layerInputBuilder{
		variable: "",
		layer:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerInputBuilder) Create() LayerInputBuilder {
	return createLayerInputBuilder()
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

	return createLayerInput(app.variable, app.layer), nil
}
