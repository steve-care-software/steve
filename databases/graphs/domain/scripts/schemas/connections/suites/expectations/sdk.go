package expectations

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewExpectationBuilder creates a new expectation builder
func NewExpectationBuilder() ExpectationBuilder {
	return createExpectationBuilder()
}

// Builder represents the expectations builder
type Builder interface {
	Create() Builder
	WithList(list []Expectation) Builder
	Now() (Expectations, error)
}

// Expectations represents expectations
type Expectations interface {
	List() []Expectation
}

// ExpectationBuilder represents the expectation builder
type ExpectationBuilder interface {
	Create() ExpectationBuilder
	WithReferences(references references.References) ExpectationBuilder
	IsFail() ExpectationBuilder
	Now() (Expectation, error)
}

// Expectation represents a suite expectation
type Expectation interface {
	References() references.References
	IsFail() bool
}
