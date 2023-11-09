package failures

// Builder represents a failure builder
type Builder interface {
	Create() Builder
	WithUsernameAlreadyExists(username string) Builder
	Now() (Failure, error)
}

// Failure represents a failure
type Failure interface {
	UsernameAlreadyExists() string
}
