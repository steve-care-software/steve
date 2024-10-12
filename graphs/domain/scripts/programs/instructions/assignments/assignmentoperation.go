package assignments

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"

type assignmentOperation struct {
	assignee   Assignee
	operator   uint8
	assignable assignables.Assignable
}

func createAssignmentOperation(assignee Assignee, operator uint8, assignable assignables.Assignable) AssignmentOperation {
	return &assignmentOperation{
		assignee:   assignee,
		operator:   operator,
		assignable: assignable,
	}
}

// Assignee returns the assignee of the operation
func (obj *assignmentOperation) Assignee() Assignee {
	return obj.assignee
}

// Operator returns the arithmetic operator of the operation
func (obj *assignmentOperation) Operator() uint8 {
	return obj.operator
}

// Assignable returns the assignable value of the operation
func (obj *assignmentOperation) Assignable() assignables.Assignable {
	return obj.assignable
}
