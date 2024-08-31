package relations

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections"
	"github.com/steve-care-software/steve/domain/relations/data/points"
)

// Builder represents the relation builder
type Builder interface {
	Create() Builder
	WithConnections(connections connections.Connections) Builder
	WithPoints(points points.Points) Builder
	Now() (Relation, error)
}

// Relation represents a relation
type Relation interface {
	Connections() connections.Connections
	HasPoints() bool
	Points() points.Points
}
