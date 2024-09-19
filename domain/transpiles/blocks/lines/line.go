package lines

import (
	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines/tokens"
)

type line struct {
	tokens tokens.Tokens
	suites suites.Suites
}

func createLine(
	tokens tokens.Tokens,
) Line {
	return createLineInternally(tokens, nil)
}

func createLineWithSuites(
	tokens tokens.Tokens,
	suites suites.Suites,
) Line {
	return createLineInternally(tokens, suites)
}

func createLineInternally(
	tokens tokens.Tokens,
	suites suites.Suites,
) Line {
	out := line{
		tokens: tokens,
		suites: suites,
	}

	return &out
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
