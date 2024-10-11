package instructions

import (
	"errors"
)

type operationBuilder struct {
	first       Assignable
	assignables OperatorAssignables
}

func createOperationBuilder() OperationBuilder {
	return &operationBuilder{
		first:       nil,
		assignables: nil,
	}
}

// Create initializes the operation builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder()
}

// WithFirst adds the first assignable to the builder
func (app *operationBuilder) WithFirst(first Assignable) OperationBuilder {
	app.first = first
	return app
}

// WithAssignables adds the operator assignables to the builder
func (app *operationBuilder) WithAssignables(assignables OperatorAssignables) OperationBuilder {
	app.assignables = assignables
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	if app.first == nil {
		return nil, errors.New("the first assignable is mandatory to build an Operation instance")
	}

	if app.assignables != nil {
		return createOperationWithAssignables(app.first, app.assignables), nil
	}

	return createOperation(app.first), nil
}
