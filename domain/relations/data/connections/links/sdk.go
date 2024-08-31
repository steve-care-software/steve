package links

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
	IsLeft() LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Name() string
	IsLeft() bool
}
