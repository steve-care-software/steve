package links

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/connections/links/contexts"
	"github.com/steve-care-software/steve/domain/hash"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	context     contexts.Context
	name        string
	isLeft      bool
	weight      float32
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		context:     nil,
		name:        "",
		isLeft:      false,
		weight:      0.0,
	}

	return &out
}

// Create initializes the builder
func (app *linkBuilder) Create() LinkBuilder {
	return createLinkBuilder(
		app.hashAdapter,
	)
}

// WithContext add context to the builder
func (app *linkBuilder) WithContext(context contexts.Context) LinkBuilder {
	app.context = context
	return app
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

// IsLeft flags the builder as left
func (app *linkBuilder) IsLeft() LinkBuilder {
	app.isLeft = true
	return app
}

// Now builds a new Link instance
func (app *linkBuilder) Now() (Link, error) {
	if app.context == nil {
		return nil, errors.New("the context is mandatory in order to build a Link instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Link instance")
	}

	if app.weight <= 0.0 {
		return nil, errors.New("the weight must be greater than 0.0 in order to build a Link instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.context.Hash().Bytes(),
		[]byte(app.name),
		[]byte(strconv.FormatFloat(float64(app.weight), 'f', 10, 32)),
	})

	if err != nil {
		return nil, err
	}

	return createLink(
		*pHash,
		app.context,
		app.name,
		app.isLeft,
		app.weight,
	), nil
}
