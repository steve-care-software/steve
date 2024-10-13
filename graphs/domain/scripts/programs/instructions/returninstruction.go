package instructions

import "github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"

type returnInstruction struct {
	assignable assignables.Assignable
}

func createReturnInstruction() ReturnInstruction {
	return createReturnInstructionInternally(nil)
}

func createReturnInstructionWithAssignable(
	assignable assignables.Assignable,
) ReturnInstruction {
	return createReturnInstructionInternally(assignable)
}

func createReturnInstructionInternally(
	assignable assignables.Assignable,
) ReturnInstruction {
	out := returnInstruction{
		assignable: assignable,
	}

	return &out
}

func (obj *returnInstruction) HasAssignable() bool {
	return obj.assignable != nil
}

func (obj *returnInstruction) Assignable() assignables.Assignable {
	return obj.assignable
}
