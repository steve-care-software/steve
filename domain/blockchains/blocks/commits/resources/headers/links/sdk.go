package links

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithContainer(container []string) Builder
	WithName(name string) Builder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Container() []string
	Name() string
}
