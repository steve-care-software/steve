package assignments

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
)

type assignmentMultipleBuilder struct {
	assignees   Assignees
	assignables assignables.Assignables
}

func createAssignmentMultipleBuilder() AssignmentMultipleBuilder {
	return &assignmentMultipleBuilder{
		assignees:   nil,
		assignables: nil,
	}
}

// Create initializes the builder
func (obj *assignmentMultipleBuilder) Create() AssignmentMultipleBuilder {
	return createAssignmentMultipleBuilder()
}

// WithAssignees sets the assignees in the builder
func (obj *assignmentMultipleBuilder) WithAssignees(assignees Assignees) AssignmentMultipleBuilder {
	obj.assignees = assignees
	return obj
}

// WithAssignables sets the assignables in the builder
func (obj *assignmentMultipleBuilder) WithAssignables(assignables assignables.Assignables) AssignmentMultipleBuilder {
	obj.assignables = assignables
	return obj
}

// Now builds a new AssignmentMultiple instance
func (obj *assignmentMultipleBuilder) Now() (AssignmentMultiple, error) {
	if obj.assignees == nil {
		return nil, errors.New("the assignees are mandatory in order to build an AssignmentMultiple instance")
	}

	if obj.assignables == nil {
		return nil, errors.New("the assignables are mandatory in order to build an AssignmentMultiple instance")
	}

	return createAssignmentMultiple(obj.assignees, obj.assignables), nil
}
