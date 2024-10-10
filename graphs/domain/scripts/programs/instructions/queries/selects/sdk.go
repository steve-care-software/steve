package selects

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the select builder
type Builder interface {
	Create() Builder
	WithExternals(externals externals.Externals) Builder
	WithCondition(condition conditions.Condition) Builder
	IsDelete() Builder
	Now() (Select, error)
}

// Select represents a select
type Select interface {
	IsDelete() bool
	Externals() externals.Externals
	HasCondition() bool
	Condition() conditions.Condition
}
