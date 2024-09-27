package suites

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
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

// SuiteBuilder represents the suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithInit(init instructions.Instructions) SuiteBuilder
	WithInput(input []byte) SuiteBuilder
	WithExpectation(expectation []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	Init() instructions.Instructions
	Input() []byte
	Expectation() []byte
}
