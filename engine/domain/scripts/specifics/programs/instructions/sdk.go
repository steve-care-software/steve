package instructions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	hashAdapter := hash.NewAdapter()
	return createInstructionBuilder(
		hashAdapter,
	)
}

// NewLoopBuilder creates a new loop builder
func NewLoopBuilder() LoopBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopBuilder(
		hashAdapter,
	)
}

// NewLoopInstructionsBuilder creates a new loop instructions builder
func NewLoopInstructionsBuilder() LoopInstructionsBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopInstructionsBuilder(
		hashAdapter,
	)
}

// NewLoopInstructionBuilder creates a new loop instruction builder
func NewLoopInstructionBuilder() LoopInstructionBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopInstructionBuilder(
		hashAdapter,
	)
}

// NewLoopHeaderBuilder creates a new loop header builder
func NewLoopHeaderBuilder() LoopHeaderBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopHeaderBuilder(
		hashAdapter,
	)
}

// NewLoopCounterBuilder creates a new loop counter builder
func NewLoopCounterBuilder() LoopCounterBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopCounterBuilder(
		hashAdapter,
	)
}

// NewLoopKeyValueBuilder creates a new loop key value builder
func NewLoopKeyValueBuilder() LoopKeyValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createLoopKeyValueBuilder(
		hashAdapter,
	)
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// Builder represents the instructions builder
type Builder interface {
	Create() Builder
	WithList(list []Instruction) Builder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment assignments.Assignment) InstructionBuilder
	WithLoop(loop Loop) InstructionBuilder
	WithCondition(condition Condition) InstructionBuilder
	WithCall(call calls.Call) InstructionBuilder
	IsReturn() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsLoop() bool
	Loop() Loop
	IsCondition() bool
	Condition() Condition
	IsCall() bool
	Call() calls.Call
	IsReturn() bool
}

// LoopBuilder represents the loop builder
type LoopBuilder interface {
	Create() LoopBuilder
	WithHeader(header LoopHeader) LoopBuilder
	WithInstructions(instructions LoopInstructions) LoopBuilder
	Now() (Loop, error)
}

// Loop represents a loop
type Loop interface {
	Hash() hash.Hash
	Header() LoopHeader
	Instructions() LoopInstructions
}

// LoopInstructionsBuilder represents the loop instructions builder
type LoopInstructionsBuilder interface {
	Create() LoopInstructionsBuilder
	WithList(list []LoopInstruction) LoopInstructionsBuilder
	Now() (LoopInstructions, error)
}

// LoopInstructions represents a loop instruction
type LoopInstructions interface {
	Hash() hash.Hash
	List() []LoopInstruction
}

// LoopInstructionBuilder represents the loop instruction builder
type LoopInstructionBuilder interface {
	Create() LoopInstructionBuilder
	WithInstruction(instruction Instruction) LoopInstructionBuilder
	IsBreak() LoopInstructionBuilder
	Now() (LoopInstruction, error)
}

// LoopInstruction represents a loop instruction
type LoopInstruction interface {
	Hash() hash.Hash
	IsInstruction() bool
	Instruction() Instruction
	IsBreak() bool
}

// LoopHeaderBuilder represents a loop header builder
type LoopHeaderBuilder interface {
	Create() LoopHeaderBuilder
	WithCounter(counter LoopCounter) LoopHeaderBuilder
	WithKeyValue(keyValue LoopKeyValue) LoopHeaderBuilder
	IsInfinite() LoopHeaderBuilder
	Now() (LoopHeader, error)
}

// LoopHeader represents a loop header
type LoopHeader interface {
	Hash() hash.Hash
	IsCounter() bool
	Counter() LoopCounter
	IsKeyValue() bool
	KeyValue() LoopKeyValue
	IsInfinite() bool
}

// LoopCounterBuilder represents the loop counter builder
type LoopCounterBuilder interface {
	Create() LoopCounterBuilder
	WithAssignment(assignment assignments.Assignment) LoopCounterBuilder
	WithOperation(operation operations.Operation) LoopCounterBuilder
	WithIncrement(increment operations.Operation) LoopCounterBuilder
	Now() (LoopCounter, error)
}

// LoopCounter represents a loop counter
type LoopCounter interface {
	Hash() hash.Hash
	Assignment() assignments.Assignment
	Operation() operations.Operation
	Increment() operations.Operation
}

type LoopKeyValueBuilder interface {
	Create() LoopKeyValueBuilder
	WithKeyname(keyname string) LoopKeyValueBuilder
	WithValueName(valueName string) LoopKeyValueBuilder
	WithOperation(operation operations.Operation) LoopKeyValueBuilder
	Now() (LoopKeyValue, error)
}

// LoopKeyValue represents a key -> value loop
type LoopKeyValue interface {
	Hash() hash.Hash
	KeyName() string
	ValueName() string
	Operation() operations.Operation
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithOperation(operation operations.Operation) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Operation() operations.Operation
	Instructions() Instructions
}
