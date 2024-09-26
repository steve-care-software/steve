package programs

import (
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/assignments"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/initializations"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

// Program represents a program
type Program interface {
	Head() heads.Head
	Input() string
	Instructions() Instructions
	HasSuites() bool
	Suites() Suites
}

// Function represents a function
type Function interface {
	Parameters() FuncParameters
	Instructions() Instructions
	HasOutput() bool
	Output() containers.Containers
	HasSuites() bool
	Suites() Suites
}

// FuncParameters represents func parameters
type FuncParameters interface {
	List() []FuncParameter
}

// FuncParameter represents a func parameter
type FuncParameter interface {
	Name() string
	Container() containers.Container
	IsMandatory() bool
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// Suite represents a test suite
type Suite interface {
	Input() []byte
	Expectation() []byte
	Instructions() Instructions
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
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
	LoopHeader() LoopHeader
	Instructions() LoopInstructions
}

// LoopInstructions represents a loop instruction
type LoopInstructions interface {
	List() []LoopInstruction
}

// LoopInstruction represents a loop instruction
type LoopInstruction interface {
	IsInstruction() bool
	Instruction() Instruction
	IsBreak() bool
}

// LoopHeader represents a loop header
type LoopHeader interface {
	IsCounter() bool
	Counter() LoopCounter
	IsKeyValue() bool
	KeyValue() LoopKeyValue
	IsInfinite() bool
}

// LoopCounter represents a loop counter
type LoopCounter interface {
	Initialization() initializations.Initialization
	Operation() operations.Operation
	Increment() operations.Operation
}

// LoopKeyValue represents a key -> value loop
type LoopKeyValue interface {
	KeyName() string
	ValueName() string
	Values() any
}

// Condition represents a condition
type Condition interface {
	Operation() operations.Operation
	Instructions() Instructions
}
