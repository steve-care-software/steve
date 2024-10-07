package permissions

import (
	"errors"
)

type builder struct {
	list []Permission
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
func (app *builder) WithList(list []Permission) Builder {
	app.list = list
	return app
}

// Now builds a new Permissions instance
func (app *builder) Now() (Permissions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Permission in order to build a Permissions instance")
	}

	return createPermissions(app.list), nil
}
