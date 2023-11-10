package executions

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithLayer(layer layers.LayerInput) Builder
	WithValues(values layers.ValueAssignments) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Layer() layers.LayerInput
	HasValues() bool
	Values() layers.ValueAssignments
}
