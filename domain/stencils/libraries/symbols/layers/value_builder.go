package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type valueBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	constant    []byte
	layer       Layer
}

func createValueBuilder(
	hashAdapter hash.Adapter,
) ValueBuilder {
	out := valueBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		constant:    nil,
		layer:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder(
		app.hashAdapter,
	)
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
	if app.constant != nil && len(app.constant) <= 0 {
		app.constant = nil
	}

	data := [][]byte{}
	if app.variable != "" {
		data = append(data, []byte(app.variable))
	}

	if app.constant != nil {
		data = append(data, app.constant)
	}

	if app.layer != nil {
		data = append(data, app.layer.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.variable != "" {
		return createValueWithVariable(*pHash, app.variable), nil
	}

	if app.constant != nil {
		return createValueWithConstant(*pHash, app.constant), nil
	}

	if app.layer != nil {
		return createValueWithLayer(*pHash, app.layer), nil
	}

	return nil, errors.New("the Value is invalid")
}
