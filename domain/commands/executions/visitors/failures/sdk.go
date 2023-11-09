package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	AdminAlreadyInitialized() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	AdminAlreadyInitialized() bool
}
