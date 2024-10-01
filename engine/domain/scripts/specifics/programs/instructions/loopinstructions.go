package instructions

import "github.com/steve-care-software/steve/engine/domain/hash"

type loopInstructions struct {
	hash hash.Hash
	list []LoopInstruction
}

func createLoopInstructions(
	hash hash.Hash,
	list []LoopInstruction,
) LoopInstructions {
	out := loopInstructions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *loopInstructions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *loopInstructions) List() []LoopInstruction {
	return obj.list
}
