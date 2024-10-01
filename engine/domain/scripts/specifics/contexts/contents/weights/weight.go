package weights

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/suites"
)

type weight struct {
	hash    hash.Hash
	name    string
	value   uint
	reverse string
	suites  suites.Suites
}

func createWeight(
	hash hash.Hash,
	name string,
	value uint,
) Weight {
	return createWeightInternally(hash, name, value, "", nil)
}

func createWeightWithReverse(
	hash hash.Hash,
	name string,
	value uint,
	reverse string,
) Weight {
	return createWeightInternally(hash, name, value, reverse, nil)
}

func createWeightWithSuites(
	hash hash.Hash,
	name string,
	value uint,
	suites suites.Suites,
) Weight {
	return createWeightInternally(hash, name, value, "", suites)
}

func createWeightWithReverseAndSuites(
	hash hash.Hash,
	name string,
	value uint,
	reverse string,
	suites suites.Suites,
) Weight {
	return createWeightInternally(hash, name, value, reverse, suites)
}

func createWeightInternally(
	hash hash.Hash,
	name string,
	value uint,
	reverse string,
	suites suites.Suites,
) Weight {
	out := weight{
		hash:    hash,
		name:    name,
		value:   value,
		reverse: reverse,
		suites:  suites,
	}

	return &out
}

// Hash returns the hash
func (obj *weight) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *weight) Name() string {
	return obj.name
}

// Value returns the name
func (obj *weight) Value() uint {
	return obj.value
}

// HasReverse returns true if there is a reverse, false otherwise
func (obj *weight) HasReverse() bool {
	return obj.reverse != ""
}

// Reverse returns the reverse, if any
func (obj *weight) Reverse() string {
	return obj.reverse
}

// HasSuites returns true if there is suites, false otherwise
func (obj *weight) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns suites, if any
func (obj *weight) Suites() suites.Suites {
	return obj.suites
}
