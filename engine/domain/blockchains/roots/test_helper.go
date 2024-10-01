package roots

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/commons/hash"
)

// NewRootForTests creates a new root for tests
func NewRootForTests(amount uint64, owner ed25519.PublicKey, commit hash.Hash) Root {
	ins, err := NewBuilder().Create().WithAmount(amount).WithOwner(owner).WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
