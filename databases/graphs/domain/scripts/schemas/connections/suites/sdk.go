package suites

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites/expectations"
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
	WithLink(link links.Link) SuiteBuilder
	WithExpectations(expectation expectations.Expectations) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Name() string
	Link() links.Link
	Expectations() expectations.Expectations
}
