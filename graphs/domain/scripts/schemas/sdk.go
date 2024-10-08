package schemas

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the schema builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithPoints(points []string) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Head() heads.Head
	Points() []string
	Connections() connections.Connections
}
