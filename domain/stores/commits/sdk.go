package commits

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/commits/modifications"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithModifications(modifications modifications.Modifications) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Modifications() modifications.Modifications
	HasParent() bool
	Parent() hash.Hash
}
