package bridges

import (
	"github.com/steve-care-software/steve/domain/graphs/connections"
	"github.com/steve-care-software/steve/domain/hash"
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
	WithConnection(connection connections.Connection) Builder
	WithWeight(weight float32) Builder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Hash() hash.Hash
	Connection() connections.Connection
	HasWeight() bool
	Weight() float32
}
