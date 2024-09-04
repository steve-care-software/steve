package roots

import (
	"github.com/steve-care-software/steve/domain/hash"
)

type root struct {
	hash   hash.Hash
	amount uint64
	owner  hash.Hash
}

func createRoot(
	hash hash.Hash,
	amount uint64,
	owner hash.Hash,
) Root {
	out := root{
		hash:   hash,
		amount: amount,
		owner:  owner,
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
func (obj *root) Owner() hash.Hash {
	return obj.owner
}
