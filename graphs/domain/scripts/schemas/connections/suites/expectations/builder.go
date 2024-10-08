package expectations

import (
	"errors"
)

type builder struct {
	list []Expectation
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
func (app *builder) WithList(list []Expectation) Builder {
	app.list = list
	return app
}

// Now builds a new Expectations instance
func (app *builder) Now() (Expectations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Expectation in order to build a Expectations instance")
	}

	return createExpectations(app.list), nil
}
