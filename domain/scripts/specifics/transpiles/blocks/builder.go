package blocks

import (
	"errors"
)

type builder struct {
	list []Block
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
func (app *builder) WithList(list []Block) Builder {
	app.list = list
	return app
}

// Now builds a new Blocks instance
func (app *builder) Now() (Blocks, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Block in order to build a Blocks instance")
	}

	return createBlocks(
		app.list,
	), nil
}
