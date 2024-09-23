package lines

import (
	"errors"

	"github.com/steve-care-software/steve/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks/lines/tokens"
)

type lineBuilder struct {
	tokens tokens.Tokens
	suites suites.Suites
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		tokens: nil,
		suites: nil,
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

// WithSuites add suites to the builder
func (app *lineBuilder) WithSuites(suites suites.Suites) LineBuilder {
	app.suites = suites
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens == nil {
		return nil, errors.New("the tokens are mandatory in order to build a Line instance")
	}

	if app.suites != nil {
		return createLineWithSuites(app.tokens, app.suites), nil
	}

	return createLine(app.tokens), nil
}
