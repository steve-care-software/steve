package targets

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the target builder
type Builder interface {
	Create() Builder
	WithConstant(constant string) Builder
	WithRule(rule string) Builder
	Now() (Target, error)
}

// Target represents a target
type Target interface {
	Hash() hash.Hash
	IsConstant() bool
	Constant() string
	IsRule() bool
	Rule() string
}
