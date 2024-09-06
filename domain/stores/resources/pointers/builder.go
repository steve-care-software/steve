package pointers

import (
	"errors"
	"fmt"
)

type builder struct {
	list []Pointer
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
func (app *builder) WithList(list []Pointer) Builder {
	app.list = list
	return app
}

// Now builds a new Pointers instance
func (app *builder) Now() (Pointers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Pointer in order to build a Pointers instance")
	}

	nextIndex := -1
	for _, onePointer := range app.list {
		index := onePointer.Index()
		if nextIndex != -1 && index != uint(nextIndex) {
			str := fmt.Sprintf("the pointer's index was expected to be %d, %d provided", nextIndex, index)
			return nil, errors.New(str)
		}

		length := onePointer.Length()
		nextIndex = int(index + length)
	}

	return createPointers(
		app.list,
	), nil
}
