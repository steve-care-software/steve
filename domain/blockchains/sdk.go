package blockchains

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
)

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithHead(head blocks.Block) Builder
	WithRoot(root roots.Root) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Hash() hash.Hash
	Head() blocks.Block
	Root() roots.Root
}
