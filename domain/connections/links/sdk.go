package links

import "github.com/steve-care-software/steve/domain/connections/links/contexts"

// NewBuilder creates a new builder
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
	WithContexts(contexts contexts.Contexts) LinkBuilder
	WithName(name string) LinkBuilder
	IsLeft() LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Name() string
	IsLeft() bool
	Contexts() contexts.Contexts
}
