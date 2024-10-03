package connections

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	return createConnectionBuilder()
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
}

// ConnectionBuilder represents the connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithHeader(header headers.Header) ConnectionBuilder
	WithLinks(links links.Links) ConnectionBuilder
	WithSuites(suites suites.Suites) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Header() headers.Header
	Links() links.Links
	HasSuites() bool
	Suites() suites.Suites
}
