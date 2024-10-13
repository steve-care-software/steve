package assignables

import (
	"errors"
)

type listMapBuilder struct {
	list Assignables
	mp   MapKeyValues
}

func createListMapBuilder() ListMapBuilder {
	return &listMapBuilder{
		list: nil,
		mp:   nil,
	}
}

// Create initializes the ListMap builder
func (app *listMapBuilder) Create() ListMapBuilder {
	return createListMapBuilder()
}

// WithList adds a list to the builder
func (app *listMapBuilder) WithList(list Assignables) ListMapBuilder {
	app.list = list
	return app
}

// WithMap adds a map to the builder
func (app *listMapBuilder) WithMap(mp MapKeyValues) ListMapBuilder {
	app.mp = mp
	return app
}

// Now builds a new ListMap instance
func (app *listMapBuilder) Now() (ListMap, error) {
	if app.list != nil {
		return createListMapWithList(app.list), nil
	}

	if app.mp != nil {
		return createListMapWithMap(app.mp), nil
	}

	return nil, errors.New("the ListMap is invalid")
}
