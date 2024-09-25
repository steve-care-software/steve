package parameters

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type parameterBuilder struct {
	hashAdapter hash.Adapter
	current     string
	local       string
}

func createPArameterBuilder(
	hashAdapter hash.Adapter,
) ParameterBuilder {
	out := parameterBuilder{
		hashAdapter: hashAdapter,
		current:     "",
		local:       "",
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createPArameterBuilder(
		app.hashAdapter,
	)
}

// WithCurrent adds a current to the builder
func (app *parameterBuilder) WithCurrent(current string) ParameterBuilder {
	app.current = current
	return app
}

// WithLocal adds a local to the builder
func (app *parameterBuilder) WithLocal(local string) ParameterBuilder {
	app.local = local
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.current == "" {
		return nil, errors.New("the current is mandatory in order to build a Parameter instance")
	}

	if app.local == "" {
		return nil, errors.New("the local is mandatory in order to build a Parameter instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.current),
		[]byte(app.local),
	})

	if err != nil {
		return nil, err
	}

	return createParameter(
		*pHash,
		app.current,
		app.local,
	), nil
}
