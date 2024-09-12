package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewContentForTests creates a new content for tests
func NewContentForTests(trx transactions.Transactions, parent hash.Hash) Content {
	ins, err := NewBuilder().Create().WithTransactions(trx).WithParent(parent).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
