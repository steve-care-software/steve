package links

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas/points"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
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
	Hash() hash.Hash
	List() []Link
}

// LinkBuilder represents the link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithOrigin(origin points.Point) LinkBuilder
	WithTarget(target points.Point) LinkBuilder
	Now() (Link, error)
}

// Link represents a point link
type Link interface {
	Hash() hash.Hash
	Origin() points.Point
	Target() points.Point
}
