package pointers

import "github.com/steve-care-software/steve/domain/stencils/pointers/symbols"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Builder represents a pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents symbol pointers
type Pointers interface {
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
	Path() []string
	Symbol() symbols.Symbol
}
