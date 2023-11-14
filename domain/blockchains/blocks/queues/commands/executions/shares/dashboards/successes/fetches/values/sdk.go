package values

import "github.com/steve-care-software/steve/domain/dashboards/stencils"

// Builder represents the fetch builder
type Builder interface {
	Create() Builder
	WithRoot(root stencils.Stencil) Builder
	WithStencils(stencils stencils.Stencils) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsRoot() bool
	Root() stencils.Stencil
	IsStencils() bool
	Stencils() stencils.Stencils
}
