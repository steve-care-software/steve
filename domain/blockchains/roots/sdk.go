package roots

import (
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// / Builder represents the root builder
type Builder interface {
	Create() Builder
	WithAmount(amount uint64) Builder
	WithOwner(owner hash.Hash) Builder
	Now() (Root, error)
}

// Root represents a root block
type Root interface {
	Hash() hash.Hash
	Amount() uint64
	Owner() hash.Hash
}
