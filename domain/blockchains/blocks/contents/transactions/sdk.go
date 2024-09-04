package transactions

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewTransactionBuilder creates a new transaction builder
func NewTransactionBuilder() TransactionBuilder {
	hashAdapter := hash.NewAdapter()
	return createTransactionBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// TransactionBuilder reprents the transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithEntry(entry entries.Entry) TransactionBuilder
	WithSignature(signature []byte) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Entry() entries.Entry
	Signature() []byte
}
