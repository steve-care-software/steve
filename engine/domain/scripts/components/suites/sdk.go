package suites

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites/expectations"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuiteBuilder(
		hashAdapter,
	)
}

// Builder represents the suites builder
type Builder interface {
	Create() Builder
	WithList(list []Suite) Builder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithOrigin(origin string) SuiteBuilder
	WithDestination(destination string) SuiteBuilder
	WithExpectations(expectations expectations.Expectations) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Name() string
	Origin() string
	Destination() string
	Expectations() expectations.Expectations
}
