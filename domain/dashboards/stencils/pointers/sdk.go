package pointers

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Builder represents the pointer builder
type Builder interface {
	Create() Builder
	WithContainer(container []string) Builder
	WithName(name string) Builder
	Now() (Pointer, error)
}

// Pointer represents a symbol pointer
type Pointer interface {
	Hash() hash.Hash
	Container() []string
	Name() string
}
