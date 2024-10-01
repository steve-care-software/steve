package constants

import "errors"

type builder struct {
	list []Constant
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Constant) Builder {
	app.list = list
	return app
}

// Now builds a new Constants instance
func (app *builder) Now() (Constants, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Constant in order to build an Constants instance")
	}

	return createOprations(
		app.list,
	), nil
}
