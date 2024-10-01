package points

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Point
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
func (app *builder) WithList(list []Point) Builder {
	app.list = list
	return app
}

// Now builds a new Points instance
func (app *builder) Now() (Points, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Point in order to build a Points instance")
	}

	data := [][]byte{}
	for _, onePoint := range app.list {
		data = append(data, onePoint.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createPoints(
		*pHash,
		app.list,
	), nil
}
