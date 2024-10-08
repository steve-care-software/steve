package links

import "errors"

type builder struct {
	list []Link
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
func (app *builder) WithList(list []Link) Builder {
	app.list = list
	return app
}

// Now builds a new Links instance
func (app *builder) Now() (Links, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Link in order to build a Links instance")
	}

	return createLinks(app.list), nil
}
