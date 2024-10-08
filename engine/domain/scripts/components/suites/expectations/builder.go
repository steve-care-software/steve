package expectations

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Expectation
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
func (app *builder) WithList(list []Expectation) Builder {
	app.list = list
	return app
}

// Now builds a new Expectations instance
func (app *builder) Now() (Expectations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Expectation in order to build a Expectations instance")
	}

	data := [][]byte{}
	for _, oneExpectation := range app.list {
		data = append(data, oneExpectation.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createExpectations(
		*pHash,
		app.list,
	), nil
}
