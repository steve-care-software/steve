package instructions

import (
	"errors"
)

type operatorAssignablesBuilder struct {
	list []OperatorAssignable
}

func createOperatorAssignablesBuilder() OperatorAssignablesBuilder {
	out := operatorAssignablesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *operatorAssignablesBuilder) Create() OperatorAssignablesBuilder {
	return createOperatorAssignablesBuilder()
}

// WithList adds a list to the builder
func (app *operatorAssignablesBuilder) WithList(list []OperatorAssignable) OperatorAssignablesBuilder {
	app.list = list
	return app
}

// Now builds a new OperatorAssignables instance
func (app *operatorAssignablesBuilder) Now() (OperatorAssignables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 OperatorAssignable in order to build a OperatorAssignables instance")
	}

	return createOperatorAssignables(app.list), nil
}
