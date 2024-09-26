package programs

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/assignments"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/initializations"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

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

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Head() heads.Head
	Input() string
	Instructions() Instructions
	HasSuites() bool
	Suites() Suites
}

// Function represents a function
type Function interface {
	Hash() hash.Hash
	Parameters() FuncParameters
	Instructions() Instructions
	HasOutput() bool
	Output() containers.Containers
	HasSuites() bool
	Suites() Suites
}

// FuncParameters represents func parameters
type FuncParameters interface {
	Hash() hash.Hash
	List() []FuncParameter
}

// FuncParameter represents a func parameter
type FuncParameter interface {
	Hash() hash.Hash
	Name() string
	Container() containers.Container
	IsMandatory() bool
}

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// Suite represents a test suite
type Suite interface {
	Hash() hash.Hash
	Init() Instructions
	Input() []byte
	Expectation() []byte
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	IsInitialization() bool
	Initialization() initializations.Initialization
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsOperation() bool
	Operation() operations.Operation
	IsLoop() bool
	Loop() Loop
	IsCondition() bool
	Condition() Condition
	IsCall() bool
	Call() calls.Call
	IsReturn() bool
}

// Loop represents a loop
type Loop interface {
	Hash() hash.Hash
	LoopHeader() LoopHeader
	Instructions() LoopInstructions
}

// LoopInstructions represents a loop instruction
type LoopInstructions interface {
	Hash() hash.Hash
	List() []LoopInstruction
}

// LoopInstruction represents a loop instruction
type LoopInstruction interface {
	Hash() hash.Hash
	IsInstruction() bool
	Instruction() Instruction
	IsBreak() bool
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
	WithInitialization(initialization initializations.Initialization) LoopCounterBuilder
	WithOperation(operation operations.Operation) LoopCounterBuilder
	WithIncrement(increment operations.Operation) LoopCounterBuilder
	Now() (LoopCounter, error)
}

// LoopCounter represents a loop counter
type LoopCounter interface {
	Hash() hash.Hash
	Initialization() initializations.Initialization
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
