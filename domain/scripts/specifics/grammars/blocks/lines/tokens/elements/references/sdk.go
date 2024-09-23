package references

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the reference builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar string) Builder
	WithBlock(block string) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Hash() hash.Hash
	Grammar() string
	Block() string
}
