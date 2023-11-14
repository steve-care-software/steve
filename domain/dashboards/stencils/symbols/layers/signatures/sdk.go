package signatures

import (
	"github.com/steve-care-software/steve/domain/dashboards/stencils/kinds"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithBytes(bytes string) Builder
	WithParams(params []string) Builder
	WithReturns(ret kinds.Kind) Builder
	WithDependencies(dependencies []string) Builder
	Now() (Signature, error)
}

// Signature represents the layer signature
type Signature interface {
	Bytes() string
	Params() []string
	Returns() kinds.Kind
	HasDependencies() bool
	Dependencies() []string
}
