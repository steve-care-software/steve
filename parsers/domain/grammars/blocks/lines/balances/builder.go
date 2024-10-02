package balances

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors"
)

type builder struct {
	lines []selectors.Selectors
}

func createBuilder() Builder {
	out := builder{
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithLines add lines to the builder
func (app *builder) WithLines(lines []selectors.Selectors) Builder {
	app.lines = lines
	return app
}

// Now builds a new Balance instance
func (app *builder) Now() (Balance, error) {
	if app.lines != nil && len(app.lines) <= 0 {
		app.lines = nil
	}

	if app.lines == nil {
		return nil, errors.New("there must be at least 1 Operations line in order to build a Balance instance")
	}

	return createBalance(
		app.lines,
	), nil
}
