package lines

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a lines builder
type Builder interface {
	Create() Builder
	WithList(list []tokens.Tokens) Builder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	Hash() hash.Hash
	List() []tokens.Tokens
}
