package schemas

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/points"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithPoints(points points.Points) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Schema, error)
}

// Schema represents the schema
type Schema interface {
	Hash() hash.Hash
	Head() heads.Head
	Points() points.Points
	Connections() connections.Connections
}
