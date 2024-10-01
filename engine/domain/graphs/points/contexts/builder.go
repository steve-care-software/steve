package contexts

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Context
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{}
	for _, oneContext := range app.list {
		data = append(data, oneContext.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createContexts(*pHash, app.list), nil
}
