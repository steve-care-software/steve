package transpiles

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs"
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// ParserAdapter represents the transpile parser adapter
type ParserAdapter interface {
	// ToTranspile takes the origin grammar and an input and create a transpile instance
	ToTranspile(origin programs.Program, input []byte) (Transpile, []byte, error)
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
	WithOrigin(origin hash.Hash) Builder
	WithTarget(target hash.Hash) Builder
	Now() (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Blocks() blocks.Blocks
	Root() string
	Origin() hash.Hash
	Target() hash.Hash
}
