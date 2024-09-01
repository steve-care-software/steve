package links

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type linkBuilder struct {
	hashAdapter hash.Adapter
	name        string
	reverse     string
}

func createLinkBuilder(
	hashAdapter hash.Adapter,
) LinkBuilder {
	out := linkBuilder{
		hashAdapter: hashAdapter,
		name:        "",
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

	data := [][]byte{
		[]byte(app.name),
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
			app.reverse,
		), nil
	}

	return createLink(
		*pHash,
		app.name,
	), nil
}
