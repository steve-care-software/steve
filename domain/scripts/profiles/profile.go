package profiles

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/domain/hash"
)

type profile struct {
	hash        hash.Hash
	version     uint
	handle      string
	name        string
	description string
	pubKey      ed25519.PublicKey
}

func createProfile(
	hash hash.Hash,
	version uint,
	handle string,
	name string,
	description string,
	pubKey ed25519.PublicKey,
) Profile {
	out := profile{
		hash:        hash,
		version:     version,
		handle:      handle,
		name:        name,
		description: description,
		pubKey:      pubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *profile) Hash() hash.Hash {
	return obj.hash
}

// Version returns the version
func (obj *profile) Version() uint {
	return obj.version
}

// Handle returns the handle
func (obj *profile) Handle() string {
	return obj.handle
}

// Name returns the name
func (obj *profile) Name() string {
	return obj.name
}

// Description returns the description
func (obj *profile) Description() string {
	return obj.description
}

// PublicKey returns the publicKey
func (obj *profile) PublicKey() ed25519.PublicKey {
	return obj.pubKey
}
