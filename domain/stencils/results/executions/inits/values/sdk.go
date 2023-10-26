package values

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"

// Builder represents a values builder
type Builder interface {
	Create() Builder
	WithList(list []Value) Builder
	Now() (Values, error)
}

// Values represents values
type Values interface {
	List() []Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithVariable(variable string) ValueBuilder
	WithBytes(bytes []byte) ValueBuilder
	WithLayer(layer layers.Layer) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Variable() string
	Content() Content
}

// Content represents a value content
type Content interface {
	IsBytes() bool
	Bytes() []byte
	IsLayer() bool
	Layer() layers.Layer
}
