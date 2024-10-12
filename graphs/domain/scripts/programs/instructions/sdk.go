package instructions

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
)

// NewForUntilClauseBuilder creates a new for until clause builder
func NewForUntilClauseBuilder() ForUntilClauseBuilder {
	return createForUntilClauseBuilder()
}

// NewKeyValueBuilder creates a new key value builderg
func NewKeyValueBuilder() ForKeyValueBuilder {
	return createForKeyValueBuilder()
}

// NewForInstructionsBuilder creates a new for instructions builder
func NewForInstructionsBuilder() ForInstructionsBuilder {
	return createForInstructionsBuilder()
}

// NewForInstructionBuilder creates a new for instruction builder
func NewForInstructionBuilder() ForInstructionBuilder {
	return createForInstructionBuilder()
}

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
	Condition() conditions.Condition
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

// ForUntilClauseBuilder represents the for until clause builder
type ForUntilClauseBuilder interface {
	Create() ForUntilClauseBuilder
	WithName(name string) ForUntilClauseBuilder
	WithValue(value uint) ForUntilClauseBuilder
	Now() (ForUntilClause, error)
}

// ForUntilClause represents the for until clause
type ForUntilClause interface {
	Name() string
	Value() uint
}

// ForKeyValueBuilder represents the for key->value builder
type ForKeyValueBuilder interface {
	Create() ForKeyValueBuilder
	WithKey(key string) ForKeyValueBuilder
	WithValue(value string) ForKeyValueBuilder
	WithIterable(iterable assignables.Iterable) ForKeyValueBuilder
	WithInstructions(instructions ForInstructions) ForKeyValueBuilder
	Now() (ForKeyValue, error)
}

// ForKeyValue represents the for key->value
type ForKeyValue interface {
	Key() string
	Value() string
	Iterable() assignables.Iterable
	Instructions() ForInstructions
}

// ForInstructionsBuilder represents the for instructions builder
type ForInstructionsBuilder interface {
	Create() ForInstructionsBuilder
	WithList(list []ForInstruction) ForInstructionsBuilder
	Now() (ForInstructions, error)
}

// ForInstructions represents the for instructions
type ForInstructions interface {
	List() []ForInstruction
}

// ForInstructionBuilder represents the for instruction builder
type ForInstructionBuilder interface {
	Create() ForInstructionBuilder
	WithInstruction(instruction Instruction) ForInstructionBuilder
	IsBreak() ForInstructionBuilder
	Now() (ForInstruction, error)
}

// ForInstruction represents the for instruction
type ForInstruction interface {
	IsBreak() bool
	IsInstruction() bool
	Instruction() Instruction
}
