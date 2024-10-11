package assignables

import (
	"errors"
)

// mapKeyValueBuilder represents the implementation of the MapKeyValueBuilder interface
type mapKeyValueBuilder struct {
	name       string
	assignable Assignable
}

// createMapKeyValueBuilder creates a new mapKeyValueBuilder instance
func createMapKeyValueBuilder() MapKeyValueBuilder {
	return &mapKeyValueBuilder{
		name:       "",
		assignable: nil,
	}
}

// Create initializes the map key-value builder
func (obj *mapKeyValueBuilder) Create() MapKeyValueBuilder {
	return createMapKeyValueBuilder()
}

// WithName adds a name (key) to the builder
func (obj *mapKeyValueBuilder) WithName(name string) MapKeyValueBuilder {
	obj.name = name
	return obj
}

// WithAssignable adds an assignable value to the builder
func (obj *mapKeyValueBuilder) WithAssignable(assignable Assignable) MapKeyValueBuilder {
	obj.assignable = assignable
	return obj
}

// Now builds a new MapKeyValue instance
func (obj *mapKeyValueBuilder) Now() (MapKeyValue, error) {
	if obj.name == "" {
		return nil, errors.New("the name (key) is mandatory in order to build a MapKeyValue instance")
	}

	if obj.assignable == nil {
		return nil, errors.New("the assignable value is mandatory in order to build a MapKeyValue instance")
	}

	return createMapKeyValue(obj.name, obj.assignable), nil
}
