package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
)

type layerBuilder struct {
	input      string
	executions Executions
	ret        return_expectations.Expectation
	params     parameters.Parameters
	suites     Suites
}

func createLayerBuilder() LayerBuilder {
	out := layerBuilder{
		input:      "",
		executions: nil,
		ret:        nil,
		params:     nil,
		suites:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerBuilder) Create() LayerBuilder {
	return createLayerBuilder()
}

// WithInput adds an input to the builder
func (app *layerBuilder) WithInput(input string) LayerBuilder {
	app.input = input
	return app
}

// WithExecutions adds an executions to the builder
func (app *layerBuilder) WithExecutions(executions Executions) LayerBuilder {
	app.executions = executions
	return app
}

// WithReturn adds a return to the builder
func (app *layerBuilder) WithReturn(ret return_expectations.Expectation) LayerBuilder {
	app.ret = ret
	return app
}

// WithParams add params to the builder
func (app *layerBuilder) WithParams(params parameters.Parameters) LayerBuilder {
	app.params = params
	return app
}

// WithSuites add suites to the builder
func (app *layerBuilder) WithSuites(suites Suites) LayerBuilder {
	app.suites = suites
	return app
}

// Now builds a new Layer instance
func (app *layerBuilder) Now() (Layer, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Layer instance")
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Layer instance")
	}

	if app.ret == nil {
		return nil, errors.New("the retunr is mandatory in order to build a Layer instance")
	}

	if app.params != nil && app.suites != nil {
		return createLayerWithParamsAndSuites(app.input, app.executions, app.ret, app.params, app.suites), nil
	}

	if app.params != nil {
		return createLayerWithParams(app.input, app.executions, app.ret, app.params), nil
	}

	if app.suites != nil {
		return createLayerWithSuites(app.input, app.executions, app.ret, app.suites), nil
	}

	return createLayer(app.input, app.executions, app.ret), nil
}
