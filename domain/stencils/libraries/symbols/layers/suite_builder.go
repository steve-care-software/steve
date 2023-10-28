package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	result_returns "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	name        string
	input       []byte
	ret         result_returns.Return
	values      ValueAssignments
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		input:       nil,
		ret:         nil,
		values:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithInput adds an input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// WithReturn adds a previous to the builder
func (app *suiteBuilder) WithReturn(ret result_returns.Return) SuiteBuilder {
	app.ret = ret
	return app
}

// WithValues add values to the builder
func (app *suiteBuilder) WithValues(values ValueAssignments) SuiteBuilder {
	app.values = values
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Suite instance")
	}

	if app.ret == nil {
		return nil, errors.New("the return is mandatory in order to build a Suite instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.input,
		app.ret.Hash().Bytes(),
	}

	if app.values != nil {
		data = append(data, app.values.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.values != nil {
		return createSuiteWithValues(*pHash, app.name, app.input, app.ret, app.values), nil
	}

	return createSuite(*pHash, app.name, app.input, app.ret), nil
}
