package kinds

import "github.com/steve-care-software/steve/domain/blockchains/hash"

// Builder represents a kind builder
type Builder interface {
	Create() Builder
	IsBytes() Builder
	IsLayer() Builder
}

// Kind represents a symbol kind
type Kind interface {
	Hash() hash.Hash
	IsBytes() bool
	IsLayer() bool
}
