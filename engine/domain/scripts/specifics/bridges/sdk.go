package bridges

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/bridges/connections"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the bridge builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithOrigin(origin string) Builder
	WithTarget(target string) Builder
	WithConnections(connections connections.Connections) Builder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Hash() hash.Hash
	Head() heads.Head
	Origin() string
	Target() string
	Connections() connections.Connections
}
