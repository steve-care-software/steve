package identities

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

// NewIdentityForTests creates a new identity for tests
func NewIdentityForTests(name string, pk ed25519.PrivateKey) Identity {
	ins, err := NewBuilder().Create().WithName(name).WithPK(pk).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIdentityWithFlagsForTests creates a new identity with flags for tests
func NewIdentityWithFlagsForTests(name string, pk ed25519.PrivateKey, flags []hash.Hash) Identity {
	ins, err := NewBuilder().Create().WithName(name).WithPK(pk).WithFlags(flags).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
