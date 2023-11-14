package symbols

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/links"
)

// Builder represents the symbols builder
type Builder interface {
	Create() Builder
	WithList(list []Symbol) Builder
	Now() (Symbols, error)
}

// Symbols represents symbols
type Symbols interface {
	Hash() hash.Hash
	List() []Symbol
}

// SymbolBuilder represents the symbol builder
type SymbolBuilder interface {
	Create() SymbolBuilder
	WithName(name string) SymbolBuilder
	WithValue(value Value) SymbolBuilder
	Now() (Symbol, error)
}

// Symbol represents a symbol
type Symbol interface {
	Hash() hash.Hash
	Name() string
	Value() Value
}

// ValueBuilder represents the value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithBytes(bytes []byte) ValueBuilder
	WithLayer(layer layers.Layer) ValueBuilder
	WithLink(link links.Link) ValueBuilder
	Now() (Value, error)
}

// Value represents the symbol value
type Value interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsLayer() layers.Layer
	IsLink() bool
	Link() links.Link
}
