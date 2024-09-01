package bridges

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/connections/links"
	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	link        links.Link
	weight      float32
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		link:        nil,
		weight:      0.0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLink adds a link to the builder
func (app *builder) WithLink(link links.Link) Builder {
	app.link = link
	return app
}

// WithWeight adds a weight to the builder
func (app *builder) WithWeight(weight float32) Builder {
	app.weight = weight
	return app
}

// Now builds a new Bridge instance
func (app *builder) Now() (Bridge, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Bridge instance")
	}

	if app.weight <= 0.0 {
		return nil, errors.New("the weight is mandatory in order to build a Bridge instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		[]byte(strconv.FormatFloat(float64(app.weight), 'f', 10, 32)),
	})

	if err != nil {
		return nil, err
	}

	return createBridge(*pHash, app.link, app.weight), nil
}
