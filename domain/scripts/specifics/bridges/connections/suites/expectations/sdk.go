package expectations

import "github.com/steve-care-software/steve/domain/hash"

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
