package connections

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
)

type connection struct {
	hash   hash.Hash
	name   string
	origin string
	target string
	suites suites.Suites
}

func createConnection(
	hash hash.Hash,
	name string,
	origin string,
	target string,
) Connection {
	return createConnectionInternally(hash, name, origin, target, nil)
}

func createConnectionWithSuites(
	hash hash.Hash,
	name string,
	origin string,
	target string,
	suites suites.Suites,
) Connection {
	return createConnectionInternally(hash, name, origin, target, suites)
}

func createConnectionInternally(
	hash hash.Hash,
	name string,
	origin string,
	target string,
	suites suites.Suites,
) Connection {
	out := connection{
		hash:   hash,
		name:   name,
		origin: origin,
		target: target,
		suites: suites,
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

// Origin returns the origin
func (obj *connection) Origin() string {
	return obj.origin
}

// Target returns the target
func (obj *connection) Target() string {
	return obj.target
}

// HasSuites returns true if there is suites, false otherwise
func (obj *connection) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *connection) Suites() suites.Suites {
	return obj.suites
}
