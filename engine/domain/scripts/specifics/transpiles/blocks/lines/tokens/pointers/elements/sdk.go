package elements

import "github.com/steve-care-software/steve/commons/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the element builder
type Builder interface {
	Create() Builder
	WithToken(token string) Builder
	WithRule(rule string) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	IsToken() bool
	Token() string
	IsRule() bool
	Rule() string
}
