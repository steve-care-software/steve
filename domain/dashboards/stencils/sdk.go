package stencils

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Builder represents the stencils builder
type Builder interface {
	Create() Builder
	WithList(list []Stencil) Builder
	Now() (Stencils, error)
}

// Stencils represents stencils
type Stencils interface {
	Hash() hash.Hash
	List() []Stencil
	Fetch(hash hash.Hash) (Stencil, error)
}

// StencilBuilder represents a stencil builder
type StencilBuilder interface {
	Create() StencilBuilder
	WithContainer(container []string) StencilBuilder
	WithDescription(description string) StencilBuilder
	WithSymbols(symbols symbols.Symbols) StencilBuilder
	WithRoot(root hash.Hash) StencilBuilder
	Now() (Stencil, error)
}

// Stencil represents a stencil
type Stencil interface {
	Hash() hash.Hash
	Container() []string
	Description() string
	Symbols() symbols.Symbols
	Root() layers.Layer
}
