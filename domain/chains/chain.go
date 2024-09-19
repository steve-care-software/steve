package chains

import (
	"github.com/steve-care-software/steve/domain/chains/nfts"
	"github.com/steve-care-software/steve/domain/hash"
)

type chain struct {
	hash    hash.Hash
	grammar nfts.NFT
	action  Action
}

func createChain(
	hash hash.Hash,
	grammar nfts.NFT,
	action Action,
) Chain {
	out := chain{
		hash:    hash,
		grammar: grammar,
		action:  action,
	}

	return &out
}

// Hash returns the hash
func (obj *chain) Hash() hash.Hash {
	return obj.hash
}

// Grammar returns the grammar
func (obj *chain) Grammar() nfts.NFT {
	return obj.grammar
}

// Action returns the action
func (obj *chain) Action() Action {
	return obj.action
}
