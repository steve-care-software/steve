package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/hash"
)

type instructionBuilder struct {
	hashAdapter hash.Adapter
	assignment  assignments.Assignment
	loop        Loop
	condition   Condition
	call        calls.Call
	isReturn    bool
}

func createInstructionBuilder(
	hashAdapter hash.Adapter,
) InstructionBuilder {
	out := instructionBuilder{
		hashAdapter: hashAdapter,
		assignment:  nil,
		loop:        nil,
		condition:   nil,
		call:        nil,
		isReturn:    false,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder(
		app.hashAdapter,
	)
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment assignments.Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithLoop adds a loop to the builder
func (app *instructionBuilder) WithLoop(loop Loop) InstructionBuilder {
	app.loop = loop
	return app
}

// WithCondition adds a condition to the builder
func (app *instructionBuilder) WithCondition(condition Condition) InstructionBuilder {
	app.condition = condition
	return app
}

// WithCall adds a call to the builder
func (app *instructionBuilder) WithCall(call calls.Call) InstructionBuilder {
	app.call = call
	return app
}

// IsReturn flags the builder as a return
func (app *instructionBuilder) IsReturn() InstructionBuilder {
	app.isReturn = true
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	data := [][]byte{}
	if app.assignment != nil {
		data = append(data, app.assignment.Hash().Bytes())
	}

	if app.loop != nil {
		data = append(data, app.loop.Hash().Bytes())
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	if app.call != nil {
		data = append(data, app.call.Hash().Bytes())
	}

	if app.isReturn {
		data = append(data, []byte("isReturn"))
	}

	if len(data) != 1 {
		return nil, errors.New("the Instruction is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.assignment != nil {
		return createInstructionWithAssignment(*pHash, app.assignment), nil
	}

	if app.loop != nil {
		return createInstructionWithLoop(*pHash, app.loop), nil
	}

	if app.condition != nil {
		return createInstructionWithCondition(*pHash, app.condition), nil
	}

	if app.call != nil {
		return createInstructionWithCall(*pHash, app.call), nil
	}

	return createInstructionWithReturn(*pHash), nil
}
