package contents

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/engine/domain/hash"
)

// NewContentForTests creates a new content for tests
func NewContentForTests(trx transactions.Transactions, parent hash.Hash, miner ed25519.PublicKey, commit hash.Hash) Content {
	ins, err := NewBuilder().Create().WithTransactions(trx).WithParent(parent).WithMiner(miner).WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
