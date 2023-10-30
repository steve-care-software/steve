package transactions

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/links"
	"github.com/steve-care-software/steve/domain/hash"
)

// Builder represents transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	List() []Transaction
}

// TransactionBuilder represents a transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithLink(link links.Link) TransactionBuilder
	WithKind(kind uint) TransactionBuilder
	WithBytes(bytes []byte) TransactionBuilder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Link() links.Link
	Kind() uint
	Bytes() []byte
}
