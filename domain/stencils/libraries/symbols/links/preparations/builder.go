package preparations

import "errors"

type builder struct {
	list []Preparation
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
func (app *builder) WithList(list []Preparation) Builder {
	app.list = list
	return app
}

// Now builds a new Preparations instance
func (app *builder) Now() (Preparations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Preparation in order to build a Preparations instance")
	}

	return createPreparations(app.list), nil
}
