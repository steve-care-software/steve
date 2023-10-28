package returns

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
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
	Hash() hash.Hash
	Output() []byte
	Kind() kinds.Kind
}
