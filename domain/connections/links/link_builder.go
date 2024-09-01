package links

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	name        string
	weight      float32
	reverse     string
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		weight:      0.0,
		reverse:     "",
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *linkBuilder) WithName(name string) LinkBuilder {
	app.name = name
	return app
}

// WithWeight adds a weight to the builder
func (app *linkBuilder) WithWeight(weight float32) LinkBuilder {
	app.weight = weight
	return app
}

// WithReverse adds a reverse to the builder
func (app *linkBuilder) WithReverse(reverse string) LinkBuilder {
	app.reverse = reverse
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Link instance")
	}

	if app.weight <= 0.0 {
		return nil, errors.New("the weight must be greater than 0.0 in order to build a Link instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.FormatFloat(float64(app.weight), 'f', 10, 32)),
	}

	if app.reverse != "" {
		data = append(data, []byte(app.reverse))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.reverse != "" {
		return createLinkWithReverse(
			*pHash,
			app.name,
			app.weight,
			app.reverse,
		), nil
	}

	return createLink(
		*pHash,
		app.name,
		app.weight,
	), nil
}
