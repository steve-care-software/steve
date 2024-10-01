package contents

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/engine/domain/hash"
)

const dataLengthTooSmallErrPattern = "(content) the data length was expected to be at least %d bytes, %d returned"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	trxAdapter := transactions.NewAdapter()
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		trxAdapter,
		hashAdapter,
		builder,
	)
}

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
	ToInstance(data []byte) (Content, []byte, error)
}

// Builder creates a new builder
type Builder interface {
	Create() Builder
	WithTransactions(trx transactions.Transactions) Builder
	WithParent(parent hash.Hash) Builder
	WithMiner(miner ed25519.PublicKey) Builder
	WithCommit(commit hash.Hash) Builder
	Now() (Content, error)
}

// Content represents a block content
type Content interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	Parent() hash.Hash
	Miner() ed25519.PublicKey
	Commit() hash.Hash
}
