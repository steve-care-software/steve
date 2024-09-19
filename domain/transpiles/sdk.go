package transpiles

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NFTAdapter represents the nft adapter
type NFTAdapter interface {
	ToNFT(ins Transpile) (nfts.NFT, error)
	ToInstance(nft nfts.NFT) (Transpile, error)
}

// Builder represents the transpile builder
type Builder interface {
	Create() Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithRoot(root string) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Blocks() blocks.Blocks
	Root() string
}
