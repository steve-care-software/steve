package links

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/assignments"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the linkd builder
type Builder interface {
	Create() Builder
	WithOrigin(origin assignments.Assignment) Builder
	WithTarget(target assignments.Assignment) Builder
	Now() (Link, error)
}

// Link represents a bridge link
type Link interface {
	Origin() assignments.Assignment
	Target() assignments.Assignment
}
