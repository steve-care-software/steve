package suites

import "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
}

// Builder represents a suites builder
type Builder interface {
	Create() Builder
	WithList(list []Suite) Builder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithInput(input []byte) SuiteBuilder
	WithOutput(output returns.Return) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents suite
type Suite interface {
	Name() string
	Input() []byte
	Output() returns.Return
}
