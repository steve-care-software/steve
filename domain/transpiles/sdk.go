package transpiles

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

// NFTAdapter represents the nft adapter
type NFTAdapter interface {
	ToNFT(ins Transpile) (nfts.NFT, error)
	ToInstance(nft nfts.NFT) (Transpile, error)
}

// Transpile represents a transpile
type Transpile interface {
	Hash() hash.Hash
	Blocks() blocks.Blocks
	Root() string
}
