package instructions

import (
	"errors"
)

type forInstructionBuilder struct {
	instruction Instruction
	isBreak     bool
}

func createForInstructionBuilder() ForInstructionBuilder {
	return &forInstructionBuilder{
		instruction: nil,
		isBreak:     false,
	}
}

// Create initializes the builder
func (obj *forInstructionBuilder) Create() ForInstructionBuilder {
	return createForInstructionBuilder()
}

// WithInstruction adds an instruction to the for instruction builder
func (obj *forInstructionBuilder) WithInstruction(instruction Instruction) ForInstructionBuilder {
	obj.instruction = instruction
	return obj
}

// IsBreak marks the for instruction as a break statement
func (obj *forInstructionBuilder) IsBreak() ForInstructionBuilder {
	obj.isBreak = true
	return obj
}

// Now builds and returns a ForInstruction instance
func (obj *forInstructionBuilder) Now() (ForInstruction, error) {
	if obj.isBreak {
		return createForInstructionWithBreak(), nil
	}

	if obj.instruction != nil {
		return createForInstructionWithInstruction(obj.instruction), nil
	}

	return nil, errors.New("the ForInstruction is invalid")
}
