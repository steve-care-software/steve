package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/points"
)

// NewInMemoryBuilder creates a new in-memory application builder
func NewInMemoryBuilder() InMemoryBuilder {
	connectionsBuilder := connections.NewBuilder()
	return createInMemoryBuilder(
		connectionsBuilder,
	)
}

// InMemoryBuilder represents the in-memory builder application
type InMemoryBuilder interface {
	Create() InMemoryBuilder
	WithConnections(connections connections.Connections) InMemoryBuilder
	WithPoints(points points.Points) InMemoryBuilder
	Now() (Application, error)
}

// Application represents a connection application
type Application interface {
	ListFrom(from uuid.UUID) (connections.Connections, error)
}
