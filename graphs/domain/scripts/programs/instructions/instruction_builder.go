package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

type instructionBuilder struct {
	singleVariableOperation assignables.SingleVariableOperation
	assignment              assignments.Assignment
	condition               conditions.Condition
	programCall             assignables.ProgramCall
	forLoop                 ForLoop
	returnInstruction       ReturnInstruction
}

func createInstructionBuilder() InstructionBuilder {
	return &instructionBuilder{
		singleVariableOperation: nil,
		assignment:              nil,
		condition:               nil,
		programCall:             nil,
		forLoop:                 nil,
		returnInstruction:       nil,
	}
}

func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

func (app *instructionBuilder) WithSingleVariableOperation(singleVariableOperation assignables.SingleVariableOperation) InstructionBuilder {
	app.singleVariableOperation = singleVariableOperation
	return app
}

func (app *instructionBuilder) WithAssignment(assignment assignments.Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

func (app *instructionBuilder) WithCondition(condition conditions.Condition) InstructionBuilder {
	app.condition = condition
	return app
}

func (app *instructionBuilder) WithProgramCall(programCall assignables.ProgramCall) InstructionBuilder {
	app.programCall = programCall
	return app
}

func (app *instructionBuilder) WithForLoop(forLoop ForLoop) InstructionBuilder {
	app.forLoop = forLoop
	return app
}

func (app *instructionBuilder) WithReturnInstruction(returnInstruction ReturnInstruction) InstructionBuilder {
	app.returnInstruction = returnInstruction
	return app
}

func (app *instructionBuilder) Now() (Instruction, error) {
	if app.singleVariableOperation != nil {
		return createInstructionWithSingleVariableOperation(app.singleVariableOperation), nil
	}

	if app.assignment != nil {
		return createInstructionWithAssignment(app.assignment), nil
	}

	if app.condition != nil {
		return createInstructionWithCondition(app.condition), nil
	}

	if app.programCall != nil {
		return createInstructionWithProgramCall(app.programCall), nil
	}

	if app.forLoop != nil {
		return createInstructionWithForLoop(app.forLoop), nil
	}

	if app.returnInstruction != nil {
		return createInstructionWithReturnInstruction(app.returnInstruction), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
