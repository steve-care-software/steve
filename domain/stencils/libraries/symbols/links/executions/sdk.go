package executions

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
	Layer() layers.LayerInput
	HasValues() bool
	Values() layers.ValueAssignments
}
