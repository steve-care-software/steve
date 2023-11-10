package suites

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	name        string
	input       []byte
	output      returns.Return
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		input:       nil,
		output:      nil,
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

// WithOutput adds an output to the builder
func (app *suiteBuilder) WithOutput(output returns.Return) SuiteBuilder {
	app.output = output
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

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a SUite instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.input,
		app.output.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(*pHash, app.name, app.input, app.output), nil
}
