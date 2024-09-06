package pointers

import "errors"

type builder struct {
	pIndex *uint
	length uint
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		length: 0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length uint) Builder {
	app.length = length
	return app
}

// Now builds a new Pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Pointer instance")
	}

	if app.length <= 0 {
		return nil, errors.New("the length is mandatory in order to build a Pointer instance")
	}

	return createPointer(*app.pIndex, app.length), nil
}
