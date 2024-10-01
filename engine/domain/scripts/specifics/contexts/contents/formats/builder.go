package formats

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Format
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
func (app *builder) WithList(list []Format) Builder {
	app.list = list
	return app
}

// Now builds a new Formats instance
func (app *builder) Now() (Formats, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Format in order to build a Formats instance")
	}

	data := [][]byte{}
	for _, oneFormat := range app.list {
		data = append(data, oneFormat.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createFormats(
		*pHash,
		app.list,
	), nil
}
