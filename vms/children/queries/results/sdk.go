package results

// Builder represents the result builder
type Builder interface {
	Create() Builder
	Now() (Result, error)
}

// Result represens results
type Result interface {
}
