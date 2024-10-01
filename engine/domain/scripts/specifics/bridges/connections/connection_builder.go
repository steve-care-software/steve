package connections

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
)

type connectionBuilder struct {
	hashAdapter hash.Adapter
	name        string
	origin      string
	target      string
	suites      suites.Suites
}

func createConnectionBuilder(
	hashAdapter hash.Adapter,
) ConnectionBuilder {
	out := connectionBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		origin:      "",
		target:      "",
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *connectionBuilder) WithName(name string) ConnectionBuilder {
	app.name = name
	return app
}

// WithOrigin adds an origin to the builder
func (app *connectionBuilder) WithOrigin(origin string) ConnectionBuilder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *connectionBuilder) WithTarget(target string) ConnectionBuilder {
	app.target = target
	return app
}

// WithSuites add suites to the builder
func (app *connectionBuilder) WithSuites(suites suites.Suites) ConnectionBuilder {
	app.suites = suites
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Connection instance")
	}

	if app.origin == "" {
		return nil, errors.New("the origin is mandatory in order to build a Connection instance")
	}

	if app.target == "" {
		return nil, errors.New("the target is mandatory in order to build a Connection instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(app.origin),
		[]byte(app.target),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createConnectionWithSuites(
			*pHash,
			app.name,
			app.origin,
			app.target,
			app.suites,
		), nil
	}

	return createConnection(
		*pHash,
		app.name,
		app.origin,
		app.target,
	), nil
}
