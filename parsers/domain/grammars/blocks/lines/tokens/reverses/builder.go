package reverses

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

type builder struct {
	hashAdapter hash.Adapter
	escape      elements.Element
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		escape:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithEscape adds an escape to the builder
func (app *builder) WithEscape(escape elements.Element) Builder {
	app.escape = escape
	return app
}

// Now builds a new Reverse instance
func (app *builder) Now() (Reverse, error) {
	data := [][]byte{
		[]byte("reverse"),
	}

	if app.escape != nil {
		data = append(data, app.escape.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.escape != nil {
		return createReverseWithEscape(*pHash, app.escape), nil
	}

	return createReverse(*pHash), nil
}
