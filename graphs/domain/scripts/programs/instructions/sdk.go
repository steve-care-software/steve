package instructions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
)

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsSingleVariableOperation() bool
	SingleVariableOperation() assignables.SingleVariableOperation
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsCondition() bool
	Condition() Condition
	IsProgramCall() bool
	ProgramCall() assignables.ProgramCall
	IsForLoop() bool
	ForLoop() ForLoop
	IsReturnInstruction() bool
	ReturnInstruction() ReturnInstruction
}

// ReturnInstruction represents a return instruction
type ReturnInstruction interface {
	HasAssignable() bool
	Assignable() assignables.Assignable
}

// ForLoop represents a for loop
type ForLoop interface {
	IsIndex() bool
	Index() ForIndex
	IsKeyValue() bool
	KeyValue() ForKeyValue
}

// ForIndex represents the for index
type ForIndex interface {
	Clause() ForUntilClause
	Instructions() ForInstructions
}

// ForUntilClause represents the for until clause
type ForUntilClause interface {
	Name() string
	Value() uint
}

// ForKeyValue represents the for key->value
type ForKeyValue interface {
	Key() string
	Value() string
	Iterable() Iterable
	Instructions() ForInstructions
}

// ForInstructions represents the for instructions
type ForInstructions interface {
	List() []ForInstruction
}

// ForInstruction represents the for instruction
type ForInstruction interface {
	IsBreak() bool
	IsInstruction() bool
	Instruction() Instruction
}

// Iterable represents an iterable
type Iterable interface {
	IsListMap() bool
	ListMap() assignables.ListMap
	IsVariable() bool
	Variable() string
}

// Condition represents a condition
type Condition interface {
	Operation() assignables.Operation
	Instructions() Instructions
}
