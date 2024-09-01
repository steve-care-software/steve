package links

import (
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewLinkBuilder creates a new link builder
func NewLinkBuilder() LinkBuilder {
	hashAdapter := hash.NewAdapter()
	return createLinkBuilder(
		hashAdapter,
	)
}

// Builder represents the links builder
type Builder interface {
	Create() Builder
	WithList(list []Link) Builder
	Now() (Links, error)
}

// Links represents links
type Links interface {
	List() []Link
}

// LinkBuilder represents the link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithName(name string) LinkBuilder
	WithReverse(reverse string) LinkBuilder
	WithWeight(weight float32) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Name() string
	Weight() float32
	Debug() string
	HasReverse() bool
	Reverse() string
}
