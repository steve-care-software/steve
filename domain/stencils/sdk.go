package stencils

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Stencils represents stencils
type Stencils interface {
	List() []Stencil
}

// Stencil represents a stencil
type Stencil interface {
	Name() string
	Description() string
	Root() hash.Hash
}
