package pointers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pointers/symbols"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	hashAdapter := hash.NewAdapter()
	return createPointerBuilder(
		hashAdapter,
	)
}

// Builder represents a pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents symbol pointers
type Pointers interface {
	Hash() hash.Hash
	List() []Pointer
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithPath(path []string) PointerBuilder
	WithSymbol(symbol symbols.Symbol) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a symbol pointer
type Pointer interface {
	Hash() hash.Hash
	Path() []string
	Symbol() symbols.Symbol
}
