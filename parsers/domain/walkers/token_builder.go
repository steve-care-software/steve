package walkers

import "errors"

type tokenBuilder struct {
	fn   ListFn
	next Walker
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		fn:   nil,
		next: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithFn adds a lift fn to the builder
func (app *tokenBuilder) WithFn(fn ListFn) TokenBuilder {
	app.fn = fn
	return app
}

// WithNext adds a next to the builder
func (app *tokenBuilder) WithNext(next Walker) TokenBuilder {
	app.next = next
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.fn == nil {
		return nil, errors.New("the ListFn is mandatory in order to build a Token instance")
	}

	if app.next != nil {
		return createTokenWithNext(app.fn, app.next), nil
	}

	return createToken(app.fn), nil
}
