package suites

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	name        string
	value       []byte
	isFail      bool
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		value:       nil,
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

// WithValue adds a value to the builder
func (app *suiteBuilder) WithValue(value []byte) SuiteBuilder {
	app.value = value
	return app
}

// IsFail flags the builder as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.value != nil && len(app.value) <= 0 {
		app.value = nil
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Suite instance")
	}

	isFail := "false"
	if app.isFail {
		isFail = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.value,
		[]byte(isFail),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(*pHash, app.name, app.value, app.isFail), nil
}
