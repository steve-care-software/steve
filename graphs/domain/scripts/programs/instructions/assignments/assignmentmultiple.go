package assignments

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"

type assignmentMultiple struct {
	assignees   Assignees
	assignables assignables.Assignables
}

func createAssignmentMultiple(assignees Assignees, assignables assignables.Assignables) AssignmentMultiple {
	return &assignmentMultiple{
		assignees:   assignees,
		assignables: assignables,
	}
}

// Assignees returns the list of assignees in the multiple assignment
func (obj *assignmentMultiple) Assignees() Assignees {
	return obj.assignees
}

// Assignables returns the list of assignables in the multiple assignment
func (obj *assignmentMultiple) Assignables() assignables.Assignables {
	return obj.assignables
}
