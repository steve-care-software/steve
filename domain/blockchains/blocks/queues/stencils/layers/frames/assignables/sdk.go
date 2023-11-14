package assignables

import "github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"

// Builder represents the assignables builder
type Builder interface {
	Create() Builder
	WithList(list []Assignable) Builder
	Now() (Assignables, error)
}

// Assignables represents assignables
type Assignables interface {
	List() []Assignable
	Compare() bool
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithLayer(layer layers.Layer) AssignableBuilder
	WithBytes(bytes []byte) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsLayer() bool
	Layer() layers.Layer
	IsBytes() bool
	Bytes() []byte
}
