package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	InstanceIsNotAdministrator() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	Name() string
	Content() Content
}

// Content represents a failure content
type Content interface {
	InstanceIsNotAdministrator() bool
}
