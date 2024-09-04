package transactions

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

type transaction struct {
	hash      hash.Hash
	entry     entries.Entry
	signature []byte
}

func createTransaction(
	hash hash.Hash,
	entry entries.Entry,
	signature []byte,
) Transaction {
	out := transaction{
		hash:      hash,
		entry:     entry,
		signature: signature,
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
