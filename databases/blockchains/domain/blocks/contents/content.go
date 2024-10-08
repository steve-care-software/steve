package contents

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/blockchains/domain/blocks/contents/transactions"
)

type content struct {
	hash   hash.Hash
	trx    transactions.Transactions
	parent hash.Hash
	miner  ed25519.PublicKey
	commit hash.Hash
}

func createContent(
	hash hash.Hash,
	trx transactions.Transactions,
	parent hash.Hash,
	miner ed25519.PublicKey,
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
func (obj *content) Miner() ed25519.PublicKey {
	return obj.miner
}

// Commit returns the commit hash
func (obj *content) Commit() hash.Hash {
	return obj.commit
}
