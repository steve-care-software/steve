package instructions

import "github.com/steve-care-software/steve/domain/hash"

type loop struct {
	hash         hash.Hash
	header       LoopHeader
	instructions LoopInstructions
}

func createLoop(
	hash hash.Hash,
	header LoopHeader,
	instructions LoopInstructions,
) Loop {
	out := loop{
		hash:         hash,
		header:       header,
		instructions: instructions,
	}

	return &out
}

// Hash returns the hash
func (obj *loop) Hash() hash.Hash {
	return obj.hash
}

// Header returns the header
func (obj *loop) Header() LoopHeader {
	return obj.header
}

// Instructions returns the instructions
func (obj *loop) Instructions() LoopInstructions {
	return obj.instructions
}
