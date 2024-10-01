package constants

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/constants/tokens"
)

type constantBuilder struct {
	name   string
	tokens tokens.Tokens
}

func createConstantBuilder() ConstantBuilder {
	out := constantBuilder{
		name:   "",
		tokens: nil,
	}

	return &out
}

// Create initializes the builder
func (app *constantBuilder) Create() ConstantBuilder {
	return createConstantBuilder()
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

// Now builds a new Constant instance
func (app *constantBuilder) Now() (Constant, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Constant instance")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens is mandatory in order to build a Constant instance")
	}

	return createConstant(
		app.name,
		app.tokens,
	), nil
}
