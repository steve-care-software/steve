package instructions

import "errors"

type constantBuilder struct {
	name  string
	value []byte
}

func createConstantBuilder() ConstantBuilder {
	out := constantBuilder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *constantBuilder) Create() ConstantBuilder {
	return createConstantBuilder()
}

// WithName adds a name to the builder
func (app *constantBuilder) WithName(name string) ConstantBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *constantBuilder) WithValue(value []byte) ConstantBuilder {
	app.value = value
	return app
}

// Now builds a new Constant instance
func (app *constantBuilder) Now() (Constant, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Constant instance")
	}

	if app.value != nil && len(app.value) <= 0 {
		app.value = nil
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Constant instance")
	}

	return createConstant(
		app.name,
		app.value,
	), nil
}
