package stencils

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Stencils represents stencils
type Stencils interface {
	List() []Stencil
	Fetch(hash hash.Hash) (Stencil, error)
}

// Stencil represents a stencil
type Stencil interface {
	Hash() hash.Hash
	Name() string
	Description() string
	Root() hash.Hash
}
