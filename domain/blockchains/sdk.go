package blockchains

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
)

// ActionFn represents an action func
type ActionFn = func() error

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithRoot(root roots.Root) Builder
	WithHead(head blocks.Block) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Hash() hash.Hash
	Root() roots.Root
	HasHead() bool
	Head() blocks.Block
}

// Repository represents a blockchain repository
type Repository interface {
	Retrieve(path string) (Blockchain, error)
}

// Service represents a blockchain service
type Service interface {
	Insert(path string, ins Blockchain) error
	Chain(original Blockchain, newBlock blocks.Block, action ActionFn) error
	Shrink(blockchain Blockchain, action ActionFn) error
}
