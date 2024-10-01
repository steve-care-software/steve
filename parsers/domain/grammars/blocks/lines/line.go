package lines

import (
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens"
)

type line struct {
	tokens  tokens.Tokens
	balance balances.Balance
}

func createLine(
	tokens tokens.Tokens,
) Line {
	return createLineInternally(tokens, nil)
}

func createLineWithBalance(
	tokens tokens.Tokens,
	balance balances.Balance,
) Line {
	return createLineInternally(tokens, balance)
}

func createLineInternally(
	tokens tokens.Tokens,
	balance balances.Balance,
) Line {
	out := line{
		tokens:  tokens,
		balance: balance,
	}

	return &out
}

// Tokens returns the tokens
func (obj *line) Tokens() tokens.Tokens {
	return obj.tokens
}

// HasBalance returns true if there is a balance, false otherwise
func (obj *line) HasBalance() bool {
	return obj.balance != nil
}

// Balance returns the balance, if any
func (obj *line) Balance() balances.Balance {
	return obj.balance
}
