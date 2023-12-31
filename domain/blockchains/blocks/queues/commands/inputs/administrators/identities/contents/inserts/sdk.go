package inserts

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithContainer(container []string) Builder
	Now() (Insert, error)
}

// Insert represents an insert command
type Insert interface {
	Name() string
	Container() []string
}
