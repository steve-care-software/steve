package transactions

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Entry() entries.Entry
	Signature() []byte
}
