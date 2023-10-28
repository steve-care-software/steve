package kinds

import "github.com/steve-care-software/steve/domain/hash"

// Builder represents a kind builder
type Builder interface {
	Create() Builder
	IsBytes() Builder
	IsLayer() Builder
	IsLink() Builder
}

// Kind represents a symbol kind
type Kind interface {
	Hash() hash.Hash
	IsBytes() bool
	IsLayer() bool
	IsLink() bool
}
