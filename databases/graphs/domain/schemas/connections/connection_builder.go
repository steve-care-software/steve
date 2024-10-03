package connections

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites"
)

type connectionBuilder struct {
	header headers.Header
	links  links.Links
	suites suites.Suites
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		header: nil,
		links:  nil,
		suites: nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// WithHeader adds an header to the builder
func (app *connectionBuilder) WithHeader(header headers.Header) ConnectionBuilder {
	app.header = header
	return app
}

// WithLinks add links to the builder
func (app *connectionBuilder) WithLinks(links links.Links) ConnectionBuilder {
	app.links = links
	return app
}

// WithSuites add suites to the builder
func (app *connectionBuilder) WithSuites(suites suites.Suites) ConnectionBuilder {
	app.suites = suites
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.header == nil {
		return nil, errors.New("the header is mandatory in order to build a Connection instance")
	}

	if app.links == nil {
		return nil, errors.New("the links is mandatory in order to build a Connection instance")
	}

	if app.suites != nil {
		return createConnectionWithSuites(
			app.header,
			app.links,
			app.suites,
		), nil
	}

	return createConnection(
		app.header,
		app.links,
	), nil
}
