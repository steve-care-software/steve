package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
)

type queryBuilder struct {
	hashAdapter hash.Adapter
	input       constantvalues.ConstantValue
	layer       LayerInput
	values      ValueAssignments
}

func createQueryBuilder(
	hashAdapter hash.Adapter,
) QueryBuilder {
	out := queryBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		layer:       nil,
		values:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryBuilder) Create() QueryBuilder {
	return createQueryBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *queryBuilder) WithInput(input constantvalues.ConstantValue) QueryBuilder {
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

	data := [][]byte{
		app.input.Hash().Bytes(),
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
		return createQueryWithValues(*pHash, app.input, app.layer, app.values), nil
	}

	return createQuery(*pHash, app.input, app.layer), nil
}
