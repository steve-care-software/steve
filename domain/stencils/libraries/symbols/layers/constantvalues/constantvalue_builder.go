package constantvalues

import "errors"

type constantValueBuilder struct {
	variable string
	constant []byte
}

func createConstantValueBuilder() ConstantValueBuilder {
	out := constantValueBuilder{
		variable: "",
		constant: nil,
	}

	return &out
}

// Create initializes the constant value builder
func (app *constantValueBuilder) Create() ConstantValueBuilder {
	return createConstantValueBuilder()
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
		return createConstantValueWithVariable(app.variable), nil
	}

	if app.constant != nil && len(app.constant) <= 0 {
		app.constant = nil
	}

	if app.constant != nil {
		return createConstantValueWithConstant(app.constant), nil
	}

	return nil, errors.New("the ConstantValue is invalid")
}
