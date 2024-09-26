package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type loopInstructionBuilder struct {
	hashAdapter hash.Adapter
	instruction Instruction
	isBreak     bool
}

func createLoopInstructionBuilder(
	hashAdapter hash.Adapter,
) LoopInstructionBuilder {
	out := loopInstructionBuilder{
		hashAdapter: hashAdapter,
		instruction: nil,
		isBreak:     false,
	}

	return &out
}

// Create initializes the builder
func (app *loopInstructionBuilder) Create() LoopInstructionBuilder {
	return createLoopInstructionBuilder(
		app.hashAdapter,
	)
}

// WithInstruction adds an instruction to the builder
func (app *loopInstructionBuilder) WithInstruction(instruction Instruction) LoopInstructionBuilder {
	app.instruction = instruction
	return app
}

// IsBreak flags the builder as a break
func (app *loopInstructionBuilder) IsBreak() LoopInstructionBuilder {
	app.isBreak = true
	return app
}

// Now builds a new LoopInstruction instance
func (app *loopInstructionBuilder) Now() (LoopInstruction, error) {
	data := [][]byte{}
	if app.instruction != nil {
		data = append(data, app.instruction.Hash().Bytes())
	}

	if app.isBreak {
		data = append(data, []byte("isBreak"))
	}

	if len(data) != 1 {
		return nil, errors.New("the LoopInstruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.instruction != nil {
		return createLoopInstructionWithInstruction(*pHash, app.instruction), nil
	}

	return createLoopInstructionWithBreak(*pHash), nil
}
