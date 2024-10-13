package instructions

type forInstruction struct {
	instruction Instruction
	isBreak     bool
}

func createForInstructionWithInstruction(instruction Instruction) ForInstruction {
	return createForInstructionInternally(instruction, false)
}

func createForInstructionWithBreak() ForInstruction {
	return createForInstructionInternally(nil, true)
}

func createForInstructionInternally(
	instruction Instruction,
	isBreak bool,
) ForInstruction {
	out := forInstruction{
		instruction: instruction,
		isBreak:     isBreak,
	}

	return &out
}

// IsBreak checks if the for instruction is a break statement
func (obj *forInstruction) IsBreak() bool {
	return obj.isBreak
}

// IsInstruction checks if the for instruction contains an instruction
func (obj *forInstruction) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction if available
func (obj *forInstruction) Instruction() Instruction {
	return obj.instruction
}
