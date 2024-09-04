package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type content struct {
	hash   hash.Hash
	trx    transactions.Transactions
	parent hash.Hash
}

func createContent(
	hash hash.Hash,
	trx transactions.Transactions,
	parent hash.Hash,
) Content {
	out := content{
		hash:   hash,
		trx:    trx,
		parent: parent,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// Transactions returns the trx
func (obj *content) Transactions() transactions.Transactions {
	return obj.trx
}

// Parent returns the parent hash
func (obj *content) Parent() hash.Hash {
	return obj.parent
}
