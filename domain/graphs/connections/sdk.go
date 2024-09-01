package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/graphs/connections/links"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
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
	List() []Connection
	Debug() string
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithLink(link links.Link) ConnectionBuilder
	From(from uuid.UUID) ConnectionBuilder
	To(to uuid.UUID) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Hash() hash.Hash
	From() uuid.UUID
	Link() links.Link
	To() uuid.UUID
	Debug() string
}
