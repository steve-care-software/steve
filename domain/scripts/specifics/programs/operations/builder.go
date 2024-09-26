package operations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Operation
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
func (app *builder) WithList(list []Operation) Builder {
	app.list = list
	return app
}

// Now builds a new Operations instance
func (app *builder) Now() (Operations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Operation in order to build a Operations instance")
	}

	data := [][]byte{}
	for _, oneOperation := range app.list {
		data = append(data, oneOperation.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createOperations(
		*pHash,
		app.list,
	), nil
}
