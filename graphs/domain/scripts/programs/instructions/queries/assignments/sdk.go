package assignments

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the assignment builder
type Builder interface {
	Create() Builder
	WithExternal(external externals.External) Builder
	WithVariable(variable string) Builder
	Now() (Assignment, error)
}

// Assignment represents a query assignment
type Assignment interface {
	External() externals.External
	Variable() string
}
