package suites

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Suite
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
func (app *builder) WithList(list []Suite) Builder {
	app.list = list
	return app
}

// Now builds a new Suites instance
func (app *builder) Now() (Suites, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Suite in order to build a Suites instance")
	}

	data := [][]byte{}
	for _, oneSuite := range app.list {
		data = append(data, oneSuite.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createSuites(*pHash, app.list), nil
}
