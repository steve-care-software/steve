package returns

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a return builder
type Builder interface {
	Create() Builder
	WithOutput(output []byte) Builder
	WithKind(kind kinds.Kind) Builder
	Now() (Return, error)
}

// Return represents a return
type Return interface {
	Output() []byte
	Kind() kinds.Kind
}
