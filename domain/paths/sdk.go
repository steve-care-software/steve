package paths

import (
	"github.com/steve-care-software/steve/domain/connections"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPathBuilder creates a new builder
func NewPathBuilder() PathBuilder {
	return createPathBuilder()
}

// Builder represents a paths builder
type Builder interface {
	Create() Builder
	WithList(list []Path) Builder
	Now() (Paths, error)
}

// Paths represents a list of paths
type Paths interface {
	List() []Path
	Successfuls() [][]connections.Connection
}

// PathBuilder represents a path builder
type PathBuilder interface {
	Create() PathBuilder
	WithPossibilities(paths Paths) PathBuilder
	WithDestination(destination connections.Connection) PathBuilder
	Now() (Path, error)
}

// Path represents a list of connections possibilities
type Path interface {
	Successfuls() [][]connections.Connection
	Possibilities() Paths
	HasDestination() bool
	Destination() connections.Connection
}
