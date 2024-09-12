package roots

import "github.com/steve-care-software/steve/domain/hash"

// NewRootForTests creates a new root for tests
func NewRootForTests(amount uint64, owner hash.Hash) Root {
	ins, err := NewBuilder().Create().WithAmount(amount).WithOwner(owner).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
