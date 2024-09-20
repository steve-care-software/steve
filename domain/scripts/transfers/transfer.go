package transfers

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/domain/hash"
)

type transfer struct {
	hash      hash.Hash
	version   uint
	amount    uint64
	publicKey ed25519.PublicKey
}

func createTransfer(
	hash hash.Hash,
	version uint,
	amount uint64,
	publicKey ed25519.PublicKey,
) Transfer {
	out := transfer{
		hash:      hash,
		version:   version,
		amount:    amount,
		publicKey: publicKey,
	}

	return &out
}

// Hash returns the hash
func (obj *transfer) Hash() hash.Hash {
	return obj.hash
}

// Version returns the version
func (obj *transfer) Version() uint {
	return obj.version
}

// Amount returns the amount
func (obj *transfer) Amount() uint64 {
	return obj.amount
}

// PublicKey returns the publicKey
func (obj *transfer) PublicKey() ed25519.PublicKey {
	return obj.publicKey
}
