package contexts

import "errors"

type builder struct {
	list []Context
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
func (app *builder) WithList(list []Context) Builder {
	app.list = list
	return app
}

// Now builds a new Contexts instance
func (app *builder) Now() (Contexts, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Context in order to build a Contexts instance")
	}

	return createContexts(app.list), nil
}
