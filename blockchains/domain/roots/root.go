package roots

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/hash"
)

type root struct {
	hash   hash.Hash
	amount uint64
	owner  ed25519.PublicKey
	commit hash.Hash
}

func createRoot(
	hash hash.Hash,
	amount uint64,
	owner ed25519.PublicKey,
	commit hash.Hash,
) Root {
	out := root{
		hash:   hash,
		amount: amount,
		owner:  owner,
		commit: commit,
	}

	return &out
}

// Hash returns the hash
func (obj *root) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *root) Amount() uint64 {
	return obj.amount
}

// Owner returns the owner
func (obj *root) Owner() ed25519.PublicKey {
	return obj.owner
}

// Commit returns the commit
func (obj *root) Commit() hash.Hash {
	return obj.commit
}
