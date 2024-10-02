package selectors

import "errors"

type builder struct {
	list []Selector
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
func (app *builder) WithList(list []Selector) Builder {
	app.list = list
	return app
}

// Now builds a new Selectors instance
func (app *builder) Now() (Selectors, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Selector in order to build an Selectors instance")
	}

	return createSelectors(
		app.list,
	), nil
}
