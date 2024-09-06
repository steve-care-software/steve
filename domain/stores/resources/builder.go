package resources

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Resource
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
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
		if nextIndex != -1 && index != uint(nextIndex) {
			str := fmt.Sprintf("the resource's (identifier: %s) pointer's index was expected to be %d, %d provided", oneResource.Identifier(), nextIndex, index)
			return nil, errors.New(str)
		}

		length := pointer.Length()
		nextIndex = int(index + length)
	}

	return createResources(
		app.list,
	), nil
}