package saves

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

type save struct {
	assignment assignments.Assignment
	condition  conditions.Condition
}

func createSave(
	assignment assignments.Assignment,
) Save {
	return createSaveInternally(assignment, nil)
}

func createSaveWithCondition(
	assignment assignments.Assignment,
	condition conditions.Condition,
) Save {
	return createSaveInternally(assignment, condition)
}

func createSaveInternally(
	assignment assignments.Assignment,
	condition conditions.Condition,
) Save {
	return &save{
		assignment: assignment,
		condition:  condition,
	}
}

// Assignment returns the assignment of the save
func (obj *save) Assignment() assignments.Assignment {
	return obj.assignment
}

// HasCondition returns true if the save has a condition
func (obj *save) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition of the save
func (obj *save) Condition() conditions.Condition {
	return obj.condition
}
