package points

import "errors"

type builder struct {
	list []Point
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
func (app *builder) WithList(list []Point) Builder {
	app.list = list
	return app
}

// Now builds a new Points instance
func (app *builder) Now() (Points, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Point in order to build a Points instance")
	}

	return createPoints(app.list), nil
}
