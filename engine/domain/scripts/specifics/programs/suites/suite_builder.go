package suites

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	init        instructions.Instructions
	input       []byte
	expectation []byte
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		init:        nil,
		input:       nil,
		expectation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithInit adds an init to the builder
func (app *suiteBuilder) WithInit(init instructions.Instructions) SuiteBuilder {
	app.init = init
	return app
}

// WithInput adds an input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// WithExpectation adds an expectation to the builder
func (app *suiteBuilder) WithExpectation(expectation []byte) SuiteBuilder {
	app.expectation = expectation
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.init == nil {
		return nil, errors.New("the init instructions is mandatory in order to build a Suite instance")
	}

	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Suite instance")
	}

	if app.expectation != nil && len(app.expectation) <= 0 {
		app.expectation = nil
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mandatory in order to build a Suite instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.init.Hash().Bytes(),
		app.input,
		app.expectation,
	})

	if err != nil {
		return nil, err
	}

	return createSuite(
		*pHash,
		app.init,
		app.input,
		app.expectation,
	), nil
}
