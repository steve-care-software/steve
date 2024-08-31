package paths

import (
	"errors"

	"github.com/steve-care-software/steve/domain/connections"
)

type builder struct {
	list []connections.Connections
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
func (app *builder) WithList(list []connections.Connections) Builder {
	app.list = list
	return app
}

// Now builds a new Paths instance
func (app *builder) Now() (Paths, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Path (list of Connections) in order to build a Paths instance")
	}

	return createPaths(app.list), nil
}
