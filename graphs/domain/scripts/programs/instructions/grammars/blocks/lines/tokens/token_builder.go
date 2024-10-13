package tokens

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/reverses"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	element     elements.Element
	cardinality cardinalities.Cardinality
	reverse     reverses.Reverse
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		element:     nil,
		cardinality: nil,
		reverse:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder(
		app.hashAdapter,
	)
}

// WithElement adds an element to the builder
func (app *tokenBuilder) WithElement(element elements.Element) TokenBuilder {
	app.element = element
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *tokenBuilder) WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder {
	app.cardinality = cardinality
	return app
}

// WithReverse adds a reverse to the builder
func (app *tokenBuilder) WithReverse(reverse reverses.Reverse) TokenBuilder {
	app.reverse = reverse
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Token instance")
	}

	data := [][]byte{
		app.element.Hash().Bytes(),
		app.cardinality.Hash().Bytes(),
	}

	if app.reverse != nil {
		data = append(data, app.reverse.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.reverse != nil {
		return createTokenWithReverse(*pHash, app.element, app.cardinality, app.reverse), nil
	}

	return createToken(*pHash, app.element, app.cardinality), nil
}
