package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/relations/data/connections/contexts"
	"github.com/steve-care-software/steve/domain/relations/data/connections/links"
)

// Connections represents connections
type Connections interface {
	List() []Connection
}

// Connection represents a connection
type Connection interface {
	Contexts() contexts.Contexts
	From() uuid.UUID
	Link() links.Link
	To() uuid.UUID
}
