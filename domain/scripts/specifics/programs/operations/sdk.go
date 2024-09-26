package operations

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/values"
)

const (
	// SingleSwordPlus represents the ++
	SingleSwordPlus (uint8) = iota

	// SingleSwordMinus represents the --
	SingleSwordMinus
)

const (
	// BitShiftLeft represents the bitshift left
	BitShiftLeft (uint8) = iota

	// BitShiftRight represents the bitshift right
	BitShiftRight
)

const (
	// StandardOperationPlus represents the +
	StandardOperationPlus (uint8) = iota

	// StandardOperationMinus represents the -
	StandardOperationMinus

	// StandardOperationMul represents the *
	StandardOperationMul

	// StandardOperationDiv represents the /
	StandardOperationDiv

	// StandardOperationMod represents the %
	StandardOperationMod

	// StandardOperationSmallerThan represents the <
	StandardOperationSmallerThan

	// StandardOperationSmallerThan represents the <=
	StandardOperationSmallerThanOrEqual

	// StandardOperationBiggerThan represents the >
	StandardOperationBiggerThan

	// StandardOperationBiggerThanOrEqual represents the >=
	StandardOperationBiggerThanOrEqual

	// StandardOperationEqual represents the ==
	StandardOperationEqual

	// StandardOperationNotEqual represents the !=
	StandardOperationNotEqual

	// StandardOperationAnd represents the &&
	StandardOperationAnd

	// StandardOperationOr represents the ||
	StandardOperationOr

	// StandardOperationNot represents the !
	StandardOperationNot

	// StandardOperationXor represents the xor
	StandardOperationXor
)

// Builder represents the operations builder
type Builder interface {
	Create() Builder
	WithList(list []Operation) Builder
	Now() (Operations, error)
}

// Operations represents operations
type Operations interface {
	Hash() hash.Hash
	List() Operation
}

// OperationBuilder represents the operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithStandardOperation(standardOperation StandardOperation) OperationBuilder
	WithSingleSword(singleSword SingleSwordOperation) OperationBuilder
	WithBitShift(bitshift BitShift) OperationBuilder
	WithValue(value values.Value) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation
// arithmetic: +, -, *, /, %
// relational: <, >, <=, >=, ==, !=
// logical: and, or, not, xor
type Operation interface {
	Hash() hash.Hash
	IsStandard() bool
	Standard() StandardOperation
	IsSingleSword() bool
	SingleSword() SingleSwordOperation
	IsBitShift() bool
	BitShift() BitShift
	IsValue() bool
	Value() values.Value
}

// BitShiftBuilder represents the bitshift builder
type BitShiftBuilder interface {
	Create() BitShiftBuilder
	WithOperation(operation Operation) BitShiftBuilder
	WithFlag(flag uint8) BitShiftBuilder
	WithAmount(amount uint8) BitShiftBuilder
	Now() (BitShift, error)
}

// BitShift represents a bitshift
type BitShift interface {
	Hash() hash.Hash
	Operation() Operation
	Flag() uint8 // <<, >>
	Amount() uint8
}

// StandardOperationBuilder represents the standard operation builder
type StandardOperationBuilder interface {
	Create() StandardOperationBuilder
	WithOperations(operations Operations) StandardOperationBuilder
	WithFlag(flag uint8) StandardOperationBuilder
	Now() (StandardOperation, error)
}

// StandardOperation represents a standard operation
type StandardOperation interface {
	Hash() hash.Hash
	Operations() Operations
	Flag() uint8
}

// SingleSwordOperationBuilder represents the single sword operation builder
type SingleSwordOperationBuilder interface {
	Create() SingleSwordOperationBuilder
	WithVariable(variable string) SingleSwordOperationBuilder
	WithFlag(flag uint8) SingleSwordOperationBuilder
	Now() (SingleSwordOperation, error)
}

// SingleSwordOperation represents a single sword operation
type SingleSwordOperation interface {
	Hash() hash.Hash
	Variable() string
	Flag() uint8 // ++, --
}
