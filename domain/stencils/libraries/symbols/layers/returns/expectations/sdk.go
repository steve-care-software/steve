package expectations

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an expectation builder
type Builder interface {
	Create() Builder
	WithVariable(ariable string) Builder
	WithKind(kind kinds.Kind) Builder
	Now() (Expectation, error)
}

// Expectation represents a return expectation
type Expectation interface {
	Hash() hash.Hash
	Variable() string
	Kind() kinds.Kind
}
