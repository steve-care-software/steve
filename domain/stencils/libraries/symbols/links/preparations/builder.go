package preparations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Preparation
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
func (app *builder) WithList(list []Preparation) Builder {
	app.list = list
	return app
}

// Now builds a new Preparations instance
func (app *builder) Now() (Preparations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Preparation in order to build a Preparations instance")
	}

	data := [][]byte{}
	for _, onePreparation := range app.list {
		data = append(data, onePreparation.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createPreparations(*pHash, app.list), nil
}
