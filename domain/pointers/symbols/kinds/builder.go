package kinds

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isBytes     bool
	isLayer     bool
	isLink      bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isBytes:     false,
		isLayer:     false,
		isLink:      false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsBytes flags the builder as bytes
func (app *builder) IsBytes() Builder {
	app.isBytes = true
	return app
}

// IsLayer flags the builder as layer
func (app *builder) IsLayer() Builder {
	app.isLayer = true
	return app
}

// IsLink flags the builder as link
func (app *builder) IsLink() Builder {
	app.isLink = true
	return app
}

// Now builds a new Kind instance
func (app *builder) Now() (Kind, error) {
	data := [][]byte{}
	if app.isBytes {
		data = append(data, []byte{0})
	}

	if app.isLayer {
		data = append(data, []byte{1})
	}

	if app.isLink {
		data = append(data, []byte{2})
	}

	if len(data) <= 0 {
		return nil, errors.New("the Kind is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isBytes {
		return createKindWithBytes(*pHash), nil
	}

	if app.isLayer {
		return createKindWithLayer(*pHash), nil
	}

	return createKindWithLink(*pHash), nil

}
