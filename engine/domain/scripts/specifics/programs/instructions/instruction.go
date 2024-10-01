package instructions

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls"
)

type instruction struct {
	hash       hash.Hash
	assignment assignments.Assignment
	loop       Loop
	condition  Condition
	call       calls.Call
	isReturn   bool
}

func createInstructionWithAssignment(
	hash hash.Hash,
	assignment assignments.Assignment,
) Instruction {
	return createInstructionInternally(
		hash,
		assignment,
		nil,
		nil,
		nil,
		false,
	)
}

func createInstructionWithLoop(
	hash hash.Hash,
	loop Loop,
) Instruction {
	return createInstructionInternally(
		hash,
		nil,
		loop,
		nil,
		nil,
		false,
	)
}

func createInstructionWithCondition(
	hash hash.Hash,
	condition Condition,
) Instruction {
	return createInstructionInternally(
		hash,
		nil,
		nil,
		condition,
		nil,
		false,
	)
}

func createInstructionWithCall(
	hash hash.Hash,
	call calls.Call,
) Instruction {
	return createInstructionInternally(
		hash,
		nil,
		nil,
		nil,
		call,
		false,
	)
}

func createInstructionWithReturn(
	hash hash.Hash,
) Instruction {
	return createInstructionInternally(
		hash,
		nil,
		nil,
		nil,
		nil,
		true,
	)
}

func createInstructionInternally(
	hash hash.Hash,
	assignment assignments.Assignment,
	loop Loop,
	condition Condition,
	call calls.Call,
	isReturn bool,
) Instruction {
	out := instruction{
		hash:       hash,
		assignment: assignment,
		loop:       loop,
		condition:  condition,
		call:       call,
		isReturn:   isReturn,
	}

	return &out
}

// Hash returns the hash
func (obj *instruction) Hash() hash.Hash {
	return obj.hash
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() assignments.Assignment {
	return obj.assignment
}

// IsLoop returns true if there is a loop, false otherwise
func (obj *instruction) IsLoop() bool {
	return obj.loop != nil
}

// Loop returns the loop, if any
func (obj *instruction) Loop() Loop {
	return obj.loop
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *instruction) Condition() Condition {
	return obj.condition
}

// IsCall returns true if there is a call, false otherwise
func (obj *instruction) IsCall() bool {
	return obj.call != nil
}

// Call returns the call, if any
func (obj *instruction) Call() calls.Call {
	return obj.call
}

// IsReturn returns true if there is a return, false otherwise
func (obj *instruction) IsReturn() bool {
	return obj.isReturn
}
