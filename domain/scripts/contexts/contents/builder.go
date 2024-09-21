package contents

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/formats"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/weights"
)

type builder struct {
	hashAdapter hash.Adapter
	formats     formats.Formats
	weights     weights.Weights
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		formats:     nil,
		weights:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFormats add a formats to the builder
func (app *builder) WithFormats(formats formats.Formats) Builder {
	app.formats = formats
	return app
}

// WithWeights add a weights to the builder
func (app *builder) WithWeights(weights weights.Weights) Builder {
	app.weights = weights
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	data := [][]byte{}
	if app.formats != nil {
		data = append(data, app.formats.Hash().Bytes())
	}

	if app.weights != nil {
		data = append(data, app.weights.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.formats != nil && app.weights != nil {
		return createContentWithFormatsAndWeights(
			*pHash,
			app.formats,
			app.weights,
		), nil
	}

	if app.formats != nil {
		return createContentWithFormats(
			*pHash,
			app.formats,
		), nil
	}

	if app.weights != nil {
		return createContentWithWeights(
			*pHash,
			app.weights,
		), nil
	}

	return nil, errors.New("the Content is invalid")
}
