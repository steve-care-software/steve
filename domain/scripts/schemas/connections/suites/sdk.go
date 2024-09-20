package suites

import "github.com/steve-care-software/steve/domain/hash"

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	Name() string
	Path() []string
	IsFail() bool
}
