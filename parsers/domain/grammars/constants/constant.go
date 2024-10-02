package constants

import "github.com/steve-care-software/steve/parsers/domain/grammars/constants/tokens"

type constant struct {
	name   string
	tokens tokens.Tokens
}

func createConstant(
	name string,
	tokens tokens.Tokens,
) Constant {
	out := constant{
		name:   name,
		tokens: tokens,
	}

	return &out
}

// Name returns the name
func (obj *constant) Name() string {
	return obj.name
}

// Tokens returns the tokens
func (obj *constant) Tokens() tokens.Tokens {
	return obj.tokens
}
