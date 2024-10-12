package assignments

import (
	"errors"
)

type builder struct {
	multiple  AssignmentMultiple
	operation AssignmentOperation
}

func createBuilder() Builder {
	return &builder{
		multiple:  nil,
		operation: nil,
	}
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithMultiple adds a multiple assignment to the builder
func (app *builder) WithMultiple(multiple AssignmentMultiple) Builder {
	app.multiple = multiple
	return app
}

// WithOperation adds an operation assignment to the builder
func (app *builder) WithOperation(operation AssignmentOperation) Builder {
	app.operation = operation
	return app
}

// Now builds and returns an Assignment instance
func (app *builder) Now() (Assignment, error) {
	if app.multiple != nil {
		return createAssignmentWithMultiple(app.multiple), nil
	}

	if app.operation != nil {
		return createAssignmentWithOperation(app.operation), nil
	}

	return nil, errors.New("the Assignment is invalid")
}
