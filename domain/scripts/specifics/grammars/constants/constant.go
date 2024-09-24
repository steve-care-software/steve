package constants

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants/tokens"
)

type constant struct {
	hash   hash.Hash
	name   string
	tokens tokens.Tokens
	suites suites.Suites
}

func createConstant(
	hash hash.Hash,
	name string,
	tokens tokens.Tokens,
) Constant {
	return createConstantInternally(
		hash,
		name,
		tokens,
		nil,
	)
}

func createConstantWithSuites(
	hash hash.Hash,
	name string,
	tokens tokens.Tokens,
	suites suites.Suites,
) Constant {
	return createConstantInternally(
		hash,
		name,
		tokens,
		suites,
	)
}

func createConstantInternally(
	hash hash.Hash,
	name string,
	tokens tokens.Tokens,
	suites suites.Suites,
) Constant {
	out := constant{
		hash:   hash,
		name:   name,
		tokens: tokens,
		suites: suites,
	}

	return &out
}

// Hash returns the hash
func (obj *constant) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *constant) Name() string {
	return obj.name
}

// Tokens returns the tokens
func (obj *constant) Tokens() tokens.Tokens {
	return obj.tokens
}

// HasSuites returns true if there is a suites, false otherwise
func (obj *constant) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *constant) Suites() suites.Suites {
	return obj.suites
}
