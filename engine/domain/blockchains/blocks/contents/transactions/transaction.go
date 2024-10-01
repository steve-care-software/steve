package transactions

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions/entries"
)

type transaction struct {
	hash      hash.Hash
	entry     entries.Entry
	signature []byte
	publicKey ed25519.PublicKey
}

func createTransaction(
	hash hash.Hash,
	entry entries.Entry,
	signature []byte,
	publicKey ed25519.PublicKey,
) Transaction {
	out := transaction{
		hash:      hash,
		entry:     entry,
		signature: signature,
		publicKey: publicKey,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Entry returns the entry
func (obj *transaction) Entry() entries.Entry {
	return obj.entry
}

// Signature returns the signature
func (obj *transaction) Signature() []byte {
	return obj.signature
}

// PublicKey returns the public key
func (obj *transaction) PublicKey() ed25519.PublicKey {
	return obj.publicKey
}
