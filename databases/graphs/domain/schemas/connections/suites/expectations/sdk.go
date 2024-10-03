package expectations

import "github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the expectation builder
type Builder interface {
	Create() Builder
	WithReferences(references references.References) Builder
	IsFail() Builder
	Now() (Expectation, error)
}

// Expectation represents a suite expectation
type Expectation interface {
	References() references.References
	IsFail() bool
}
