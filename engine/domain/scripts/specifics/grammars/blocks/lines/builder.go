package lines

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines/tokens"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []tokens.Tokens
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
func (app *builder) WithList(list []tokens.Tokens) Builder {
	app.list = list
	return app
}

// Now builds a new Lines instance
func (app *builder) Now() (Lines, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Lines instance")
	}

	data := [][]byte{}
	for _, oneToken := range app.list {
		data = append(data, oneToken.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createLines(
		*pHash,
		app.list,
	), nil
}
