package lines

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens"
)

type line struct {
	tokens tokens.Tokens
}

func createLine(
	tokens tokens.Tokens,
) Line {
	return createLineInternally(tokens)
}

func createLineInternally(
	tokens tokens.Tokens,
) Line {
	out := line{
		tokens: tokens,
	}

	return &out
}

// Tokens returns the tokens
func (obj *line) Tokens() tokens.Tokens {
	return obj.tokens
}
