package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type content struct {
	hash   hash.Hash
	trx    transactions.Transactions
	parent hash.Hash
	miner  hash.Hash
	commit hash.Hash
}

func createContent(
	hash hash.Hash,
	trx transactions.Transactions,
	parent hash.Hash,
	miner hash.Hash,
	commit hash.Hash,
) Content {
	out := content{
		hash:   hash,
		trx:    trx,
		parent: parent,
		miner:  miner,
		commit: commit,
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

// Miner returns the miner hash
func (obj *content) Miner() hash.Hash {
	return obj.miner
}

// Commit returns the commit hash
func (obj *content) Commit() hash.Hash {
	return obj.commit
}
