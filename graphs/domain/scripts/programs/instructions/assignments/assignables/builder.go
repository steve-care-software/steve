package assignables

import (
	"errors"
)

type builder struct {
	list []Assignable
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
func (app *builder) WithList(list []Assignable) Builder {
	app.list = list
	return app
}

// Now builds a new Assignables instance
func (app *builder) Now() (Assignables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Assignable in order to build a Assignables instance")
	}

	return createAssignables(app.list), nil
}
