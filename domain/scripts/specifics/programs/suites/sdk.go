package suites

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
)

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	Init() instructions.Instructions
	Input() []byte
	Expectation() []byte
}
