package functions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/functions/parameters"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/suites"
)

type functionBuilder struct {
	hashAdapter  hash.Adapter
	parameters   parameters.Parameters
	instructions instructions.Instructions
	output       containers.Containers
	suites       suites.Suites
}

func createFunctionBuilder(
	hashAdapter hash.Adapter,
) FunctionBuilder {
	out := functionBuilder{
		hashAdapter:  hashAdapter,
		parameters:   nil,
		instructions: nil,
		output:       nil,
		suites:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *functionBuilder) Create() FunctionBuilder {
	return createFunctionBuilder(
		app.hashAdapter,
	)
}

// WithParameters add parameters to the builder
func (app *functionBuilder) WithParameters(parameters parameters.Parameters) FunctionBuilder {
	app.parameters = parameters
	return app
}

// WithInstructions add instructions to the builder
func (app *functionBuilder) WithInstructions(instructions instructions.Instructions) FunctionBuilder {
	app.instructions = instructions
	return app
}

// WithOutput add output to the builder
func (app *functionBuilder) WithOutput(output containers.Containers) FunctionBuilder {
	app.output = output
	return app
}

// WithSuites add suites to the builder
func (app *functionBuilder) WithSuites(suites suites.Suites) FunctionBuilder {
	app.suites = suites
	return app
}

// Now builds a new Function instance
func (app *functionBuilder) Now() (Function, error) {
	if app.parameters == nil {
		return nil, errors.New("the parameters is mandatory in order to build a Function instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Function instance")
	}

	data := [][]byte{
		app.parameters.Hash().Bytes(),
		app.instructions.Hash().Bytes(),
	}

	if app.output != nil {
		data = append(data, app.output.Hash().Bytes())
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.output != nil && app.suites != nil {
		return createFunctionWithOutputAndSuites(
			*pHash,
			app.parameters,
			app.instructions,
			app.output,
			app.suites,
		), nil
	}

	if app.output != nil {
		return createFunctionWithOutput(
			*pHash,
			app.parameters,
			app.instructions,
			app.output,
		), nil
	}

	if app.suites != nil {
		return createFunctionWithSuites(
			*pHash,
			app.parameters,
			app.instructions,
			app.suites,
		), nil
	}

	return createFunction(
		*pHash,
		app.parameters,
		app.instructions,
	), nil
}
