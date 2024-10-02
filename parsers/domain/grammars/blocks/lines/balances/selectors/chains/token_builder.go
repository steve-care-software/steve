package chains

import "errors"

type tokenBuilder struct {
	pIndex  *uint
	element Element
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		pIndex:  nil,
		element: nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithIndex adds an index to the builder
func (app *tokenBuilder) WithIndex(index uint) TokenBuilder {
	app.pIndex = &index
	return app
}

// WithElement adds an element to the builder
func (app *tokenBuilder) WithElement(element Element) TokenBuilder {
	app.element = element
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Token instance")
	}

	if app.element != nil {
		return createTokenWithElement(
			*app.pIndex,
			app.element,
		), nil
	}

	return createToken(
		*app.pIndex,
	), nil
}
