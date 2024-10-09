package saves

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

// Save represents a save
type Save interface {
	Assignment() assignments.Assignment
	HasCondition() bool
	Condition() conditions.Condition
}
