package weights

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Weight
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
func (app *builder) WithList(list []Weight) Builder {
	app.list = list
	return app
}

// Now builds a new Weights instance
func (app *builder) Now() (Weights, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Weight in order to build a Weights instance")
	}

	data := [][]byte{}
	for _, oneWeight := range app.list {
		data = append(data, oneWeight.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createWeights(
		*pHash,
		app.list,
	), nil
}
