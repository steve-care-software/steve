package instructions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

type instruction struct {
	singleVariableOperation assignables.SingleVariableOperation
	assignment              assignments.Assignment
	condition               conditions.Condition
	programCall             assignables.ProgramCall
	forLoop                 ForLoop
	returnInstruction       ReturnInstruction
}

func createInstructionWithSingleVariableOperation(
	singleVariableOperation assignables.SingleVariableOperation,
) Instruction {
	return createInstructionInternally(
		singleVariableOperation,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithAssignment(
	assignment assignments.Assignment,
) Instruction {
	return createInstructionInternally(
		nil,
		assignment,
		nil,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithCondition(
	condition conditions.Condition,
) Instruction {
	return createInstructionInternally(
		nil,
		nil,
		condition,
		nil,
		nil,
		nil,
	)
}

func createInstructionWithProgramCall(
	programCall assignables.ProgramCall,
) Instruction {
	return createInstructionInternally(
		nil,
		nil,
		nil,
		programCall,
		nil,
		nil,
	)
}

func createInstructionWithForLoop(
	forLoop ForLoop,
) Instruction {
	return createInstructionInternally(
		nil,
		nil,
		nil,
		nil,
		forLoop,
		nil,
	)
}

func createInstructionWithReturnInstruction(
	returnInstruction ReturnInstruction,
) Instruction {
	return createInstructionInternally(
		nil,
		nil,
		nil,
		nil,
		nil,
		returnInstruction,
	)
}

func createInstructionInternally(
	singleVariableOperation assignables.SingleVariableOperation,
	assignment assignments.Assignment,
	condition conditions.Condition,
	programCall assignables.ProgramCall,
	forLoop ForLoop,
	returnInstruction ReturnInstruction,
) Instruction {
	out := instruction{
		singleVariableOperation: singleVariableOperation,
		assignment:              assignment,
		condition:               condition,
		programCall:             programCall,
		forLoop:                 forLoop,
		returnInstruction:       returnInstruction,
	}

	return &out
}

func (obj *instruction) IsSingleVariableOperation() bool {
	return obj.singleVariableOperation != nil
}

func (obj *instruction) SingleVariableOperation() assignables.SingleVariableOperation {
	return obj.singleVariableOperation
}

func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

func (obj *instruction) Assignment() assignments.Assignment {
	return obj.assignment
}

func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

func (obj *instruction) Condition() conditions.Condition {
	return obj.condition
}

func (obj *instruction) IsProgramCall() bool {
	return obj.programCall != nil
}

func (obj *instruction) ProgramCall() assignables.ProgramCall {
	return obj.programCall
}

func (obj *instruction) IsForLoop() bool {
	return obj.forLoop != nil
}

func (obj *instruction) ForLoop() ForLoop {
	return obj.forLoop
}

func (obj *instruction) IsReturnInstruction() bool {
	return obj.returnInstruction != nil
}

func (obj *instruction) ReturnInstruction() ReturnInstruction {
	return obj.returnInstruction
}
