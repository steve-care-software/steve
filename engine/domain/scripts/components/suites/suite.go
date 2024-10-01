package suites

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites/expectations"
)

type suite struct {
	hash         hash.Hash
	name         string
	origin       string
	destination  string
	expectations expectations.Expectations
}

func createSuite(
	hash hash.Hash,
	name string,
	origin string,
	destination string,
	expectations expectations.Expectations,
) Suite {
	out := suite{
		hash:         hash,
		name:         name,
		origin:       origin,
		destination:  destination,
		expectations: expectations,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Origin returns the origin
func (obj *suite) Origin() string {
	return obj.origin
}

// Destination returns the destination
func (obj *suite) Destination() string {
	return obj.destination
}

// Expectations returns the expectations
func (obj *suite) Expectations() expectations.Expectations {
	return obj.expectations
}
