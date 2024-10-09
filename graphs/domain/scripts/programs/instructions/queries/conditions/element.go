package conditions

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"

// element represents the implementation of the Element interface
type element struct {
	assignment assignments.Assignment
	condition  Condition
}

func createElement(assignment assignments.Assignment, condition Condition) Element {
	return &element{
		assignment: assignment,
		condition:  condition,
	}
}

// IsAssignment returns true if the element contains an assignment
func (obj *element) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment of the element
func (obj *element) Assignment() assignments.Assignment {
	return obj.assignment
}

// IsCondition returns true if the element contains a condition
func (obj *element) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition of the element
func (obj *element) Condition() Condition {
	return obj.condition
}
