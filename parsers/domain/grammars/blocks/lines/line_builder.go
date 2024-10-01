package lines

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens"
)

type lineBuilder struct {
	tokens tokens.Tokens
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		tokens: nil,
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

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Line instance")
	}

	return createLine(app.tokens), nil
}
