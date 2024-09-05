package identities

import (
	"crypto"

	"github.com/steve-care-software/steve/domain/hash"
)

type identity struct {
	name  string
	pk    crypto.PrivateKey
	flags []hash.Hash
}

func createIdentity(
	name string,
	pk crypto.PrivateKey,
) Identity {
	return createIdentityInternally(name, pk, nil)
}

func createIdentityWithFlags(
	name string,
	pk crypto.PrivateKey,
	flags []hash.Hash,
) Identity {
	return createIdentityInternally(name, pk, flags)
}

func createIdentityInternally(
	name string,
	pk crypto.PrivateKey,
	flags []hash.Hash,
) Identity {
	out := identity{
		name:  name,
		pk:    pk,
		flags: flags,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// PK returns the pk
func (obj *identity) PK() crypto.PrivateKey {
	return obj.pk
}

// HasFlags returns true if there is flags, false otherwise
func (obj *identity) HasFlags() bool {
	return obj.flags != nil
}

// Flags returns the flags, if any
func (obj *identity) Flags() []hash.Hash {
	return obj.flags
}
