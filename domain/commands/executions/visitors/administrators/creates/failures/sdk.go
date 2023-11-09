package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	AdminAlreadyExists() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	AdminAlreadyExists() bool
}
