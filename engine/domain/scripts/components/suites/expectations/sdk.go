package expectations

import "github.com/steve-care-software/steve/commons/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewExpectationBuilder creates a new expectation builder
func NewExpectationBuilder() ExpectationBuilder {
	hashAdapter := hash.NewAdapter()
	return createExpectationBuilder(
		hashAdapter,
	)
}

// Builder represents the expectations builder
type Builder interface {
	Create() Builder
	WithList(list []Expectation) Builder
	Now() (Expectations, error)
}

// Expectations represents expectations
type Expectations interface {
	Hash() hash.Hash
	List() []Expectation
}

// ExpectationBuilder represents the expectation builder
type ExpectationBuilder interface {
	Create() ExpectationBuilder
	WithPath(path []string) ExpectationBuilder
	IsFail() ExpectationBuilder
	Now() (Expectation, error)
}

// Expectation represents an expectation
type Expectation interface {
	Hash() hash.Hash
	Path() []string
	IsFail() bool
}
