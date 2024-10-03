package headers

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers/names"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an header struct
type Builder interface {
	Create() Builder
	WithName(name names.Name) Builder
	WithReverse(reverse names.Name) Builder
	Now() (Header, error)
}

// Header represents an header
type Header interface {
	Name() names.Name
	HasReverse() bool
	Reverse() names.Name
}
