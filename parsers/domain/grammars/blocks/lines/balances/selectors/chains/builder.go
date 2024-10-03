package chains

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type builder struct {
	element elements.Element
	token   Token
}

func createBuilder() Builder {
	out := builder{
		element: nil,
		token:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element elements.Element) Builder {
	app.element = element
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token Token) Builder {
	app.token = token
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Chain instance")
	}

	if app.token != nil {
		return createChainWithToken(
			app.element,
			app.token,
		), nil
	}

	return createChain(
		app.element,
	), nil
}
