package transactions

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/engine/domain/hash"
)

const dataLengthTooSmallErrPattern = "(transaction) the data length was expected to be at least %d bytes, %d returned"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	entryAdapter := entries.NewAdapter()
	builder := NewBuilder()
	transactionBuilder := NewTransactionBuilder()
	return createAdapter(
		entryAdapter,
		builder,
		transactionBuilder,
	)
}

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

// Adapter represents the transactions adapter
type Adapter interface {
	InstancesToBytes(ins Transactions) ([]byte, error)
	BytesToInstances(data []byte) (Transactions, []byte, error)
	InstanceToBytes(ins Transaction) ([]byte, error)
	BytesToInstance(data []byte) (Transaction, []byte, error)
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
	WithPublicKey(pubKey ed25519.PublicKey) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Entry() entries.Entry
	Signature() []byte
	PublicKey() ed25519.PublicKey
}
