package suites

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	name        string
	input       []byte
	isFail      bool
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		input:       nil,
		isFail:      false,
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

// WithInput adds a input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// IsFail flags the suite as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	isFail := "false"
	if app.isFail {
		isFail = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.input,
		[]byte(isFail),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(*pHash, app.name, app.input, app.isFail), nil
}
