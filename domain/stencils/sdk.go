package stencils

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/resolutions"
)

// Stencils represents stencils
type Stencils interface {
	List() []Stencil
}

// Stencil represents a stencil
type Stencil interface {
	Name() string
	Description() string
	Resolution() resolutions.Resolution
	Root() layers.Layer
}
