package assignables

import (
	"errors"
)

type mapKeyValuesBuilder struct {
	list []MapKeyValue
}

func createMapKeyValuesBuilder() MapKeyValuesBuilder {
	out := mapKeyValuesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *mapKeyValuesBuilder) Create() MapKeyValuesBuilder {
	return createMapKeyValuesBuilder()
}

// WithList adds a list to the builder
func (app *mapKeyValuesBuilder) WithList(list []MapKeyValue) MapKeyValuesBuilder {
	app.list = list
	return app
}

// Now builds a new MapKeyValues instance
func (app *mapKeyValuesBuilder) Now() (MapKeyValues, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 MapKeyValue in order to build a MapKeyValues instance")
	}

	return createMapKeyValues(app.list), nil
}
