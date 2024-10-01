package chains

import (
	"errors"

	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type builder struct {
	element       elements.Element
	pTokenIndex   *uint
	pElementIndex *uint
	next          Chain
}

func createBuilder() Builder {
	out := builder{
		element:       nil,
		pTokenIndex:   nil,
		pElementIndex: nil,
		next:          nil,
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

// WithTokenIndex adds a tokenIndex to the builder
func (app *builder) WithTokenIndex(tokenIndex uint) Builder {
	app.pTokenIndex = &tokenIndex
	return app
}

// WithElementIndex adds an elementIndex to the builder
func (app *builder) WithElementIndex(elementIndex uint) Builder {
	app.pElementIndex = &elementIndex
	return app
}

// WithNext adds a next to the builder
func (app *builder) WithNext(next Chain) Builder {
	app.next = next
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Chain instance")
	}

	if app.pTokenIndex == nil {
		return nil, errors.New("the tokenIndex is mandatory in order to build a Chain instance")
	}

	if app.pElementIndex == nil {
		return nil, errors.New("the elementIndex is mandatory in order to build a Chain instance")
	}

	if app.next != nil {
		return createChainWithNext(
			app.element,
			*app.pTokenIndex,
			*app.pElementIndex,
			app.next,
		), nil
	}

	return createChain(
		app.element,
		*app.pTokenIndex,
		*app.pElementIndex,
	), nil
}
