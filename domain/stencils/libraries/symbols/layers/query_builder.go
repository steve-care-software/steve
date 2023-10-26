package layers

import (
	"errors"
)

type queryBuilder struct {
	input  ConstantValue
	layer  LayerInput
	values ValueAssignments
}

func createQueryBuilder() QueryBuilder {
	out := queryBuilder{
		input:  nil,
		layer:  nil,
		values: nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryBuilder) Create() QueryBuilder {
	return createQueryBuilder()
}

// WithInput adds an input to the builder
func (app *queryBuilder) WithInput(input ConstantValue) QueryBuilder {
	app.input = input
	return app
}

// WithLayer adds a layer to the builder
func (app *queryBuilder) WithLayer(layer LayerInput) QueryBuilder {
	app.layer = layer
	return app
}

// WithValues add values to the builder
func (app *queryBuilder) WithValues(values ValueAssignments) QueryBuilder {
	app.values = values
	return app
}

// Now builds a new Query instance
func (app *queryBuilder) Now() (Query, error) {
	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Query instance")
	}

	if app.layer == nil {
		return nil, errors.New("the layer input is mandatory in order to build a Query instance")
	}

	if app.values != nil {
		return createQueryWithValues(app.input, app.layer, app.values), nil
	}

	return createQuery(app.input, app.layer), nil
}
