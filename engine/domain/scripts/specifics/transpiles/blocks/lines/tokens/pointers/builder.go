package pointers

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines/tokens/pointers/elements"
)

type builder struct {
	hashAdapter hash.Adapter
	element     elements.Element
	pIndex      *uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		element:     nil,
		pIndex:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.element.Hash().Bytes(),
		[]byte(strconv.Itoa(int(*app.pIndex))),
	})

	if err != nil {
		return nil, err
	}

	return createPointer(
		*pHash,
		app.element,
		*app.pIndex,
	), nil
}
