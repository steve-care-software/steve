package blockchains

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Hash() hash.Hash
	Head() blocks.Block
	Root() roots.Root
}
