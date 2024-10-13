package tokens

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants/tokens/elements"
	"github.com/steve-care-software/steve/hash"
)

type tokenBuilder struct {
	hashAdapter hash.Adapter
	element     elements.Element
	occurences  uint
}

func createTokenBuilder(
	hashAdapter hash.Adapter,
) TokenBuilder {
	out := tokenBuilder{
		hashAdapter: hashAdapter,
		element:     nil,
		occurences:  0,
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

// WithOccurences add occurences to the builder
func (app *tokenBuilder) WithOccurences(occurences uint) TokenBuilder {
	app.occurences = occurences
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.occurences <= 0 {
		return nil, errors.New("the occurences is mandatory in order to build a Token instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.element.Hash().Bytes(),
		[]byte(strconv.Itoa(int(app.occurences))),
	})

	if err != nil {
		return nil, err
	}

	return createToken(
		*pHash,
		app.element,
		app.occurences,
	), nil
}
