package layers

import "errors"

type constantValuesBuilder struct {
	list []ConstantValue
}

func createConstantValuesBuilder() ConstantValuesBuilder {
	out := constantValuesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the constantValuesBuilder
func (app *constantValuesBuilder) Create() ConstantValuesBuilder {
	return createConstantValuesBuilder()
}

// WithList adds a list to the constantValuesBuilder
func (app *constantValuesBuilder) WithList(list []ConstantValue) ConstantValuesBuilder {
	app.list = list
	return app
}

// Now builds a new ConstantValues instance
func (app *constantValuesBuilder) Now() (ConstantValues, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ConstantValue in order to build a ConstantValues instance")
	}

	return createConstantValues(app.list), nil
}
