package constants

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants/tokens"
)

// Constants represents constants
type Constants interface {
	Hash() hash.Hash
	List() []Constant
}

// Constant represents a constant
type Constant interface {
	Hash() hash.Hash
	Name() string
	Tokens() tokens.Tokens
	HasSuites() bool
	Suites() suites.Suites
}
