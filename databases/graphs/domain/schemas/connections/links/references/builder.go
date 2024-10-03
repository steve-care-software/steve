package references

import (
	"errors"
)

type builder struct {
	list []Reference
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
func (app *builder) WithList(list []Reference) Builder {
	app.list = list
	return app
}

// Now builds a new References instance
func (app *builder) Now() (References, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Reference in order to build a References instance")
	}

	return createReferences(app.list), nil
}
