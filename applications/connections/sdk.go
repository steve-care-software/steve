package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/connections"
)

// Application represents a connection application
type Application interface {
	ListFrom(from uuid.UUID) (connections.Connections, error)
	ListTo(to uuid.UUID) (connections.Connections, error)
	Retrieve(from uuid.UUID, to uuid.UUID) (connections.Connection, error)
	Insert(connection connections.Connection) error
	Delete(identifier uuid.UUID) error
}
