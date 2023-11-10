package constantvalues

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

type constantValueBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	constant    []byte
}

func createConstantValueBuilder(
	hashAdapter hash.Adapter,
) ConstantValueBuilder {
	out := constantValueBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		constant:    nil,
	}

	return &out
}

// Create initializes the constant value builder
func (app *constantValueBuilder) Create() ConstantValueBuilder {
	return createConstantValueBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *constantValueBuilder) WithVariable(variable string) ConstantValueBuilder {
	app.variable = variable
	return app
}

// WithConstant adds a constant to the builder
func (app *constantValueBuilder) WithConstant(constant []byte) ConstantValueBuilder {
	app.constant = constant
	return app
}

// Now builds a new ConstantValue instance
func (app *constantValueBuilder) Now() (ConstantValue, error) {
	if app.variable != "" {
		pHash, err := app.hashAdapter.FromString(app.variable)
		if err != nil {
			return nil, err
		}

		return createConstantValueWithVariable(*pHash, app.variable), nil
	}

	if app.constant != nil && len(app.constant) <= 0 {
		app.constant = nil
	}

	if app.constant != nil {
		pHash, err := app.hashAdapter.FromBytes(app.constant)
		if err != nil {
			return nil, err
		}

		return createConstantValueWithConstant(*pHash, app.constant), nil
	}

	return nil, errors.New("the ConstantValue is invalid")
}
