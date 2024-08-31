package paths

import (
	"github.com/steve-care-software/steve/domain/connections"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a paths builder
type Builder interface {
	Create() Builder
	WithList(list []connections.Connections) Builder
	Now() (Paths, error)
}

// Paths represents a list of connections path
type Paths interface {
	List() []connections.Connections
}
