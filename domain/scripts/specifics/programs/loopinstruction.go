package programs

import "github.com/steve-care-software/steve/domain/hash"

type loopInstruction struct {
	hash        hash.Hash
	instruction Instruction
	isBreak     bool
}

func createLoopInstructionWithInstruction(
	hash hash.Hash,
	instruction Instruction,
) LoopInstruction {
	return createLoopInstructionInternally(hash, instruction, false)
}

func createLoopInstructionWithBreak(
	hash hash.Hash,
) LoopInstruction {
	return createLoopInstructionInternally(hash, nil, true)
}

func createLoopInstructionInternally(
	hash hash.Hash,
	instruction Instruction,
	isBreak bool,
) LoopInstruction {
	out := loopInstruction{
		hash:        hash,
		instruction: instruction,
		isBreak:     isBreak,
	}

	return &out
}

// Hash returns the hash
func (obj *loopInstruction) Hash() hash.Hash {
	return obj.hash
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *loopInstruction) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *loopInstruction) Instruction() Instruction {
	return obj.instruction
}

// IsBreak return true if there is a break, false otherwise
func (obj *loopInstruction) IsBreak() bool {
	return obj.isBreak
}
