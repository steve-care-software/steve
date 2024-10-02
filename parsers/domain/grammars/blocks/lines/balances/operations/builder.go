package operations

import "errors"

type builder struct {
	list []Operation
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
func (app *builder) WithList(list []Operation) Builder {
	app.list = list
	return app
}

// Now builds a new Operations instance
func (app *builder) Now() (Operations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Operation in order to build an Operations instance")
	}

	return createOperations(
		app.list,
	), nil
}
