package pointers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers/elements"
)

type builder struct {
	element elements.Element
	pIndex  *uint
}

func createBuilder() Builder {
	out := builder{
		element: nil,
		pIndex:  nil,
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

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Pointer instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Pointer instance")
	}

	return createPointer(
		app.element,
		*app.pIndex,
	), nil
}
