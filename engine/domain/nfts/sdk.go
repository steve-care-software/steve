package nfts

import (
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new nfts builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewNFTBuilder creates a new nft builder
func NewNFTBuilder() NFTBuilder {
	hashAdapter := hash.NewAdapter()
	return createNFTBuilder(
		hashAdapter,
	)
}

// Adapter represents the nft adapter
type Adapter interface {
	ToNFT(bytes []byte) (NFT, error)
	ToInstance(ins NFT) ([]byte, error)
}

// Builder represents an nfts builder
type Builder interface {
	Create() Builder
	WithList(list []NFT) Builder
	WithName(name string) Builder
	Now() (NFTs, error)
}

// NFTs represents nfts
type NFTs interface {
	Hash() hash.Hash
	List() []NFT
	Fetch(name string) (NFT, error)
	HasName() bool
	Name() string
}

// NFTBuilder represents an nft builder
type NFTBuilder interface {
	Create() NFTBuilder
	WithByte(byte byte) NFTBuilder
	WithNFTs(nfts NFTs) NFTBuilder
	WithRecursive(recursive uint) NFTBuilder
	Now() (NFT, error)
}

// NFT represents an nft
type NFT interface {
	Hash() hash.Hash
	Fetch(name string) (NFT, error)
	IsByte() bool
	Byte() *byte
	IsNFTs() bool
	NFTs() NFTs
	IsRecursive() bool
	Recursive() *uint
}
