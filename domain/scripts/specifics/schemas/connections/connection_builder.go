package connections

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections/links"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections/suites"
)

type connectionBuilder struct {
	hashAdapter hash.Adapter
	name        string
	links       links.Links
	reverse     string
	suites      suites.Suites
}

func createConnectionBuilder(
	hashAdapter hash.Adapter,
) ConnectionBuilder {
	out := connectionBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		links:       nil,
		reverse:     "",
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

// WithLinks adds links to the builder
func (app *connectionBuilder) WithLinks(links links.Links) ConnectionBuilder {
	app.links = links
	return app
}

// WithReverse adds reverse to the builder
func (app *connectionBuilder) WithReverse(reverse string) ConnectionBuilder {
	app.reverse = reverse
	return app
}

// WithSuites adds suites to the builder
func (app *connectionBuilder) WithSuites(suites suites.Suites) ConnectionBuilder {
	app.suites = suites
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Connection instance")
	}

	if app.links == nil {
		return nil, errors.New("the links is mandatory in order to build a Connection instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.links.Hash().Bytes(),
	}

	if app.reverse != "" {
		data = append(data, []byte(app.reverse))
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.reverse != "" && app.suites != nil {
		return createConnectionWithReverseAndSuites(
			*pHash,
			app.name,
			app.links,
			app.reverse,
			app.suites,
		), nil
	}

	if app.reverse != "" {
		return createConnectionWithReverse(
			*pHash,
			app.name,
			app.links,
			app.reverse,
		), nil
	}

	if app.suites != nil {
		return createConnectionWithSuites(
			*pHash,
			app.name,
			app.links,
			app.suites,
		), nil
	}

	return createConnection(
		*pHash,
		app.name,
		app.links,
	), nil
}
