package routes

import "errors"

type builder struct {
	list []Route
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
func (app *builder) WithList(list []Route) Builder {
	app.list = list
	return app
}

// Now builds a new Routes instance
func (app *builder) Now() (Routes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Route in order to build a Routes instance")
	}

	return createRoutes(app.list), nil
}
