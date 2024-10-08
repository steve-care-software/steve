package roles

import "github.com/steve-care-software/steve/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a role builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithName(name string) Builder
	WithInsert(insert []hash.Hash) Builder
	WithDelete(del []hash.Hash) Builder
	Now() (Role, error)
}

// Role represents the role
type Role interface {
	Hash() hash.Hash
	Version() uint
	Name() string
	HasInsert() bool
	Insert() []hash.Hash
	HasDelete() bool
	Delete() []hash.Hash
}
