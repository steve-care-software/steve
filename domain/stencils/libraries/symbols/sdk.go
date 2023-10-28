package symbols

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSymbolBuilder creates a new symbol builder
func NewSymbolBuilder() SymbolBuilder {
	return createSymbolBuilder()
}

// Builder represents a symbols builder
type Builder interface {
	Create() Builder
	WithList(list []Symbol) Builder
	Now() (Symbols, error)
}

// Symbols represents symbols
type Symbols interface {
	List() []Symbol
	Fetch(name string) (Symbol, error)
}

// SymbolBuilder represents a symbol builder
type SymbolBuilder interface {
	Create() SymbolBuilder
	WithBytes(bytes []byte) SymbolBuilder
	WithLayer(layer layers.Layer) SymbolBuilder
	WithLink(link links.Link) SymbolBuilder
	Now() (Symbol, error)
}

// Symbol represents a symbol
type Symbol interface {
	IsBytes() bool
	Bytes() []byte
	IsLayer() bool
	Layer() layers.Layer
	IsLink() bool
	Link() links.Link
}

// RepositoryBuilder represents a reposiotry builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents the symbol repository
type Repository interface {
	Exists(container []string, hash hash.Hash) (bool, error)
	Retrieve(container []string, hash hash.Hash) (Symbol, error)
}

// Service represents a symbol service
type Service interface {
	Insert(context uint, container []string, symbol Symbol) error
}
