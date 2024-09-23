package connections

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/suites"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConnectionBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	Hash() hash.Hash
	List() []Connection
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithName(name string) ConnectionBuilder
	WithOrigin(origin string) ConnectionBuilder
	WithTarget(target string) ConnectionBuilder
	WithSuites(suites suites.Suites) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Hash() hash.Hash
	Name() string
	Origin() string
	Target() string
	HasSuites() bool
	Suites() suites.Suites
}
