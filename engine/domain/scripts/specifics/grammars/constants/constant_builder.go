package constants

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/constants/tokens"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type constantBuilder struct {
	hashAdapter hash.Adapter
	name        string
	tokens      tokens.Tokens
	suites      suites.Suites
}

func createConstantBuilder(
	hashAdapter hash.Adapter,
) ConstantBuilder {
	out := constantBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		tokens:      nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *constantBuilder) Create() ConstantBuilder {
	return createConstantBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *constantBuilder) WithName(name string) ConstantBuilder {
	app.name = name
	return app
}

// WithTokens add tokens to the builder
func (app *constantBuilder) WithTokens(tokens tokens.Tokens) ConstantBuilder {
	app.tokens = tokens
	return app
}

// WithSuites add suites to the builder
func (app *constantBuilder) WithSuites(suites suites.Suites) ConstantBuilder {
	app.suites = suites
	return app
}

// Now builds a new Constant instance
func (app *constantBuilder) Now() (Constant, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Constant instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens is mandatory in order to build a Constant instance")
	}

	data := [][]byte{
		[]byte(app.name),
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
		return createConstantWithSuites(
			*pHash,
			app.name,
			app.tokens,
			app.suites,
		), nil
	}

	return createConstant(
		*pHash,
		app.name,
		app.tokens,
	), nil
}
