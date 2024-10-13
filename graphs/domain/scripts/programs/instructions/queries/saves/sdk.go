package saves

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the save builder
type Builder interface {
	Create() Builder
	WithAssignment(assignment assignments.Assignment) Builder
	WithCondition(condition conditions.Condition) Builder
	Now() (Save, error)
}

// Save represents a save
type Save interface {
	Assignment() assignments.Assignment
	HasCondition() bool
	Condition() conditions.Condition
}
