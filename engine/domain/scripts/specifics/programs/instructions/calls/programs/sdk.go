package programs

import (
	"github.com/steve-care-software/steve/commons/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithInput(input string) Builder
	Now() (Program, error)
}

// Program represents a program call
type Program interface {
	Hash() hash.Hash
	Name() string
	Input() string
}
