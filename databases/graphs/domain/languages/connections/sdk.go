package connections

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/suites"
)

// Connections represents connections
type Connections interface {
	List() []Connection
}

// Connection represents a connection
type Connection interface {
	Header() headers.Header
	Links() links.Links
	HasSuites() bool
	Suites() suites.Suites
}
