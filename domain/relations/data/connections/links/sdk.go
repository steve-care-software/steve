package links

import "github.com/steve-care-software/steve/domain/relations/data/connections/links/contexts"

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
	HasContexts() bool
	Contexts() contexts.Contexts
}
