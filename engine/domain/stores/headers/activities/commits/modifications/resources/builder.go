package resources

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Resource
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
func (app *builder) WithList(list []Resource) Builder {
	app.list = list
	return app
}

// Now builds a new Resources instance
func (app *builder) Now() (Resources, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Resource in order to build a Resources instance")
	}

	nextIndex := -1
	for _, oneResource := range app.list {
		pointer := oneResource.Pointer()
		index := pointer.Index()
		if nextIndex != -1 && index < uint(nextIndex) {
			str := fmt.Sprintf("the resource's (identifier: %s) pointer's index was expected to be at least %d, %d provided", oneResource.Identifier(), nextIndex, index)
			return nil, errors.New(str)
		}

		length := pointer.Length()
		nextIndex = int(index + length)
	}

	data := [][]byte{}
	for _, oneResource := range app.list {
		data = append(data, oneResource.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createResources(
		*pHash,
		app.list,
	), nil
}
