package lines

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type lineBuilder struct {
	hashAdapter hash.Adapter
	tokens      tokens.Tokens
	suites      suites.Suites
}

func createLineBuilder(
	hashAdapter hash.Adapter,
) LineBuilder {
	out := lineBuilder{
		hashAdapter: hashAdapter,
		tokens:      nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder(
		app.hashAdapter,
	)
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

	data := [][]byte{
		app.tokens.Hash().Bytes(),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createLineWithSuites(*pHash, app.tokens, app.suites), nil
	}

	return createLine(*pHash, app.tokens), nil
}
