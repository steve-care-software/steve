package lines

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens"
)

type lineBuilder struct {
	tokens  tokens.Tokens
	balance balances.Balance
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		tokens:  nil,
		balance: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithTokens add tokens to the builder
func (app *lineBuilder) WithTokens(tokens tokens.Tokens) LineBuilder {
	app.tokens = tokens
	return app
}

// WithBalance adds a balance to the builder
func (app *lineBuilder) WithBalance(balance balances.Balance) LineBuilder {
	app.balance = balance
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Line instance")
	}

	if app.balance != nil {
		return createLineWithBalance(app.tokens, app.balance), nil
	}

	return createLine(app.tokens), nil
}
