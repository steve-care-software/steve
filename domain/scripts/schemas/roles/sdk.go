package roles

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the role builder
type Builder interface {
	Create() Builder
	WithRead(read []string) Builder
	WithWrite(write []string) Builder
	WithReview(review []string) Builder
	Now() (Role, error)
}

// Role represents role
type Role interface {
	Hash() hash.Hash
	HasRead() bool
	Read() []string
	HasWrite() bool
	Write() []string
	HasReview() bool
	Review() []string
}
