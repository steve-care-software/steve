package assignments

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
)

type assignmentOperationBuilder struct {
	assignee   Assignee
	pOperator  *uint8
	assignable assignables.Assignable
}

func createAssignmentOperationBuilder() AssignmentOperationBuilder {
	return &assignmentOperationBuilder{
		assignee:   nil,
		pOperator:  nil,
		assignable: nil,
	}
}

// Create initializes the builder
func (app *assignmentOperationBuilder) Create() AssignmentOperationBuilder {
	return createAssignmentOperationBuilder()
}

// WithAssignee sets the assignee in the builder
func (app *assignmentOperationBuilder) WithAssignee(assignee Assignee) AssignmentOperationBuilder {
	app.assignee = assignee
	return app
}

// WithOperator sets the operator in the builder
func (app *assignmentOperationBuilder) WithOperator(operator uint8) AssignmentOperationBuilder {
	app.pOperator = &operator
	return app
}

// WithAssignable sets the assignable in the builder
func (app *assignmentOperationBuilder) WithAssignable(assignable assignables.Assignable) AssignmentOperationBuilder {
	app.assignable = assignable
	return app
}

// Now builds a new AssignmentOperation instance
func (app *assignmentOperationBuilder) Now() (AssignmentOperation, error) {
	if app.assignee == nil {
		return nil, errors.New("the assignee is mandatory in order to build an AssignmentOperation instance")
	}

	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory in order to build an AssignmentOperation instance")
	}

	if app.pOperator == nil {
		return nil, errors.New("the operator is mandatory in order to build an AssignmentOperation instance")
	}

	return createAssignmentOperation(app.assignee, *app.pOperator, app.assignable), nil
}
