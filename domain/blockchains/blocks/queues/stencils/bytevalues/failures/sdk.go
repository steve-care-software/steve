package failures

// Builder represents the builder
type Builder interface {
	Create() Builder
	CouldNotFetchVariable() Builder
	VariableIsNotBytes() Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	CouldNotFetchVariable() bool
	VariableIsNotBytes() bool
}
