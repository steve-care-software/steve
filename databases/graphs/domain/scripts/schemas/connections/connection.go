package connections

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites"
)

type connection struct {
	header headers.Header
	links  links.Links
	suites suites.Suites
}

func createConnection(
	header headers.Header,
	links links.Links,
) Connection {
	return createConnectionInternally(header, links, nil)
}

func createConnectionWithSuites(
	header headers.Header,
	links links.Links,
	suites suites.Suites,
) Connection {
	return createConnectionInternally(header, links, suites)
}

func createConnectionInternally(
	header headers.Header,
	links links.Links,
	suites suites.Suites,
) Connection {
	out := connection{
		header: header,
		links:  links,
		suites: suites,
	}

	return &out
}

// Header returns the header
func (obj *connection) Header() headers.Header {
	return obj.header
}

// Links returns the links
func (obj *connection) Links() links.Links {
	return obj.links
}

// HasSuites returns true if there is a suites, false otherwise
func (obj *connection) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *connection) Suites() suites.Suites {
	return obj.suites
}
