package suites

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites/expectations"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
}

// Builder represents the suites builder
type Builder interface {
	Create() Builder
	WithList(list []Suite) Builder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteBuilder represents the suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithReference(reference references.Reference) SuiteBuilder
	WithExpectation(expectation expectations.Expectation) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Name() string
	Reference() references.Reference
	Expectation() expectations.Expectation
}
