package paths

import "errors"

type builder struct {
	list []Path
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
func (app *builder) WithList(list []Path) Builder {
	app.list = list
	return app
}

// Now builds a new Paths instance
func (app *builder) Now() (Paths, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Path in order to build a Paths instance")
	}

	return createPaths(app.list), nil
}
