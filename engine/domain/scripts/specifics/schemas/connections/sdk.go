package connections

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas/connections/links"
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

// Builder represents the connections builder
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

// ConnectionBuilder represents the connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithName(name string) ConnectionBuilder
	WithLinks(links links.Links) ConnectionBuilder
	WithReverse(reverse string) ConnectionBuilder
	WithSuites(suites suites.Suites) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Hash() hash.Hash
	Name() string
	Links() links.Links
	HasReverse() bool
	Reverse() string
	HasSuites() bool
	Suites() suites.Suites
}
