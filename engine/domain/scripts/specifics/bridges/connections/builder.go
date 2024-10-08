package connections

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Connection
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
func (app *builder) WithList(list []Connection) Builder {
	app.list = list
	return app
}

// Now builds a new Connections instance
func (app *builder) Now() (Connections, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Connection in order to build a Connections instance")
	}

	data := [][]byte{}
	for _, oneConnection := range app.list {
		data = append(data, oneConnection.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createConnections(
		*pHash,
		app.list,
	), nil
}
