package layers

import "errors"

type valueBuilder struct {
	variable string
	constant []byte
	layer    Layer
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		variable: "",
		constant: nil,
		layer:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithVariable adds a variable to the builder
func (app *valueBuilder) WithVariable(variable string) ValueBuilder {
	app.variable = variable
	return app
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant []byte) ValueBuilder {
	app.constant = constant
	return app
}

// WithLayer adds a layer to the builder
func (app *valueBuilder) WithLayer(layer Layer) ValueBuilder {
	app.layer = layer
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.variable != "" {
		return createValueWithVariable(app.variable), nil
	}

	if app.constant != nil && len(app.constant) <= 0 {
		app.constant = nil
	}

	if app.constant != nil {
		return createValueWithConstant(app.constant), nil
	}

	if app.layer != nil {
		return createValueWithLayer(app.layer), nil
	}

	return nil, errors.New("the Value is invalid")
}
