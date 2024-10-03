package links

import "github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewLinkBuilder creates a new link builder
func NewLinkBuilder() LinkBuilder {
	return createLinkBuilder()
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
	WithOrigin(origin references.Reference) LinkBuilder
	WithTarget(target references.Reference) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Origin() references.Reference
	Target() references.Reference
}
