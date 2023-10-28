package kinds

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a kind builder
type Builder interface {
	Create() Builder
	IsBytes() Builder
	IsLayer() Builder
	IsLink() Builder
	Now() (Kind, error)
}

// Kind represents a symbol kind
type Kind interface {
	Hash() hash.Hash
	IsBytes() bool
	IsLayer() bool
	IsLink() bool
}
