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
func (app *forInstructionBuilder) Create() ForInstructionBuilder {
	return createForInstructionBuilder()
}

// WithInstruction adds an instruction to the for instruction builder
func (app *forInstructionBuilder) WithInstruction(instruction Instruction) ForInstructionBuilder {
	app.instruction = instruction
	return app
}

// IsBreak marks the for instruction as a break statement
func (app *forInstructionBuilder) IsBreak() ForInstructionBuilder {
	app.isBreak = true
	return app
}

// Now builds and returns a ForInstruction instance
func (app *forInstructionBuilder) Now() (ForInstruction, error) {
	if app.isBreak {
		return createForInstructionWithBreak(), nil
	}

	if app.instruction != nil {
		return createForInstructionWithInstruction(app.instruction), nil
	}

	return nil, errors.New("the ForInstruction is invalid")
}
