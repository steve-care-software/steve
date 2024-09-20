package connections

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/domain/scripts/schemas/connections/suites"
)

// Connections represents connections
type Connections interface {
	Hash() hash.Hash
	List() []Connection
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
