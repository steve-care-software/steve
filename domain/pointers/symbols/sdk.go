package symbols

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pointers/symbols/kinds"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a symbol builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithKind(kind kinds.Kind) Builder
	Now() (Symbol, error)
}

// Symbol represents a symbol
type Symbol interface {
	Hash() hash.Hash
	Name() string
	Kind() kinds.Kind
}
