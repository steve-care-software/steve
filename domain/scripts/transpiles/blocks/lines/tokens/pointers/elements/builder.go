package elements

import (
	"errors"
)

type builder struct {
	token string
	rule  string
}

func createBuilder() Builder {
	out := builder{
		token: "",
		rule:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token string) Builder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	if app.token != "" {
		return createElementWithToken(app.token), nil
	}

	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	return nil, errors.New("the Element is invalid")
}
