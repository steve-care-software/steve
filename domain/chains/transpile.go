package chains

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

type transpile struct {
	hash   hash.Hash
	bridge nfts.NFT
	target nfts.NFT
	next   Chain
}

func createTranspile(
	hash hash.Hash,
	bridge nfts.NFT,
	target nfts.NFT,
) Transpile {
	return createTranspileInternally(hash, bridge, target, nil)
}

func createTranspileWithNext(
	hash hash.Hash,
	bridge nfts.NFT,
	target nfts.NFT,
	next Chain,
) Transpile {
	return createTranspileInternally(hash, bridge, target, next)
}

func createTranspileInternally(
	hash hash.Hash,
	bridge nfts.NFT,
	target nfts.NFT,
	next Chain,
) Transpile {
	out := transpile{
		hash:   hash,
		bridge: bridge,
		target: target,
		next:   next,
	}

	return &out
}

// Hash returns the hash
func (obj *transpile) Hash() hash.Hash {
	return obj.hash
}

// Bridge returns the bridge
func (obj *transpile) Bridge() nfts.NFT {
	return obj.bridge
}

// Target returns the target
func (obj *transpile) Target() nfts.NFT {
	return obj.target
}

// HasNext returns true if there is a next, false otherwise
func (obj *transpile) HasNext() bool {
	return obj.next != nil
}

// Next returns the next chain, if any
func (obj *transpile) Next() Chain {
	return obj.next
}
