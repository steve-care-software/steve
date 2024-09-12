package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the content adapter
type Adapter interface {
	ToBytes(ins Content) ([]byte, error)
	ToInstance(data []byte) (Content, error)
}

// Builder creates a new builder
type Builder interface {
	Create() Builder
	WithTransactions(trx transactions.Transactions) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Content, error)
}

// Content represents a block content
type Content interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	Parent() hash.Hash
}
