package pointers

import "errors"

type pointerBuilder struct {
	pIndex *uint
	length uint
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		pIndex: nil,
		length: 0,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithIndex adds an index to the builder
func (app *pointerBuilder) WithIndex(index uint) PointerBuilder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the builder
func (app *pointerBuilder) WithLength(length uint) PointerBuilder {
	app.length = length
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.pIndex != nil {
		return nil, errors.New("the index is mandatory in order to build a Pointer instance")
	}

	if app.length <= 0 {
		return nil, errors.New("the length is mandatory in order to build a Pointer instance")
	}

	return createPointer(*app.pIndex, app.length), nil
}
