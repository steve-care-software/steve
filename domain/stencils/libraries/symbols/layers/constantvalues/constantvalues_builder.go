package constantvalues

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

type constantValuesBuilder struct {
	hashAdapter hash.Adapter
	list        []ConstantValue
}

func createConstantValuesBuilder(
	hashAdapter hash.Adapter,
) ConstantValuesBuilder {
	out := constantValuesBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the constantValuesBuilder
func (app *constantValuesBuilder) Create() ConstantValuesBuilder {
	return createConstantValuesBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the constantValuesBuilder
func (app *constantValuesBuilder) WithList(list []ConstantValue) ConstantValuesBuilder {
	app.list = list
	return app
}

// Now builds a new ConstantValues instance
func (app *constantValuesBuilder) Now() (ConstantValues, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ConstantValue in order to build a ConstantValues instance")
	}

	bytes := [][]byte{}
	for _, oneValue := range app.list {
		bytes = append(bytes, oneValue.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	return createConstantValues(*pHash, app.list), nil
}
