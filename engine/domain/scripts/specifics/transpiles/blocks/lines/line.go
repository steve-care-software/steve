package lines

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type line struct {
	hash   hash.Hash
	tokens tokens.Tokens
	suites suites.Suites
}

func createLine(
	hash hash.Hash,
	tokens tokens.Tokens,
) Line {
	return createLineInternally(hash, tokens, nil)
}

func createLineWithSuites(
	hash hash.Hash,
	tokens tokens.Tokens,
	suites suites.Suites,
) Line {
	return createLineInternally(hash, tokens, suites)
}

func createLineInternally(
	hash hash.Hash,
	tokens tokens.Tokens,
	suites suites.Suites,
) Line {
	out := line{
		hash:   hash,
		tokens: tokens,
		suites: suites,
	}

	return &out
}

// Hash returns the hash
func (obj *line) Hash() hash.Hash {
	return obj.hash
}

// Tokens returns the tokens
func (obj *line) Tokens() tokens.Tokens {
	return obj.tokens
}

// HasSuites returns true if there is suites, false otherwise
func (obj *line) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *line) Suites() suites.Suites {
	return obj.suites
}
