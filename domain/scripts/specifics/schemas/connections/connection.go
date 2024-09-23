package connections

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/suites"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections/links"
)

type connection struct {
	hash    hash.Hash
	name    string
	links   links.Links
	reverse string
	suites  suites.Suites
}

func createConnection(
	hash hash.Hash,
	name string,
	links links.Links,
) Connection {
	return createConnectionInternally(
		hash,
		name,
		links,
		"",
		nil,
	)
}

func createConnectionWithReverse(
	hash hash.Hash,
	name string,
	links links.Links,
	reverse string,
) Connection {
	return createConnectionInternally(
		hash,
		name,
		links,
		reverse,
		nil,
	)
}

func createConnectionWithSuites(
	hash hash.Hash,
	name string,
	links links.Links,
	suites suites.Suites,
) Connection {
	return createConnectionInternally(
		hash,
		name,
		links,
		"",
		suites,
	)
}

func createConnectionWithReverseAndSuites(
	hash hash.Hash,
	name string,
	links links.Links,
	reverse string,
	suites suites.Suites,
) Connection {
	return createConnectionInternally(
		hash,
		name,
		links,
		reverse,
		suites,
	)
}

func createConnectionInternally(
	hash hash.Hash,
	name string,
	links links.Links,
	reverse string,
	suites suites.Suites,
) Connection {
	out := connection{
		hash:    hash,
		name:    name,
		links:   links,
		reverse: reverse,
		suites:  suites,
	}

	return &out
}

// Hash returns the hash
func (obj *connection) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *connection) Name() string {
	return obj.name
}

// Links returns the links
func (obj *connection) Links() links.Links {
	return obj.links
}

// HasReverse returns true if reverse, false otherwise
func (obj *connection) HasReverse() bool {
	return obj.reverse != ""
}

// Reverse returns the reverse, if any
func (obj *connection) Reverse() string {
	return obj.reverse
}

// HasSuites returns true if suites, false otherwise
func (obj *connection) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *connection) Suites() suites.Suites {
	return obj.suites
}
