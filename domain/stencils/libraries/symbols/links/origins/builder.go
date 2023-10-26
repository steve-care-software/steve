package origins

import "errors"

type builder struct {
	list []Origin
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
func (app *builder) WithList(list []Origin) Builder {
	app.list = list
	return app
}

// Now builds a new Origins instance
func (app *builder) Now() (Origins, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Origin in order to build a Origins instance")
	}

	return createOrigins(app.list), nil
}
