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
	// StandardPlus represents the +
	StandardPlus (uint8) = iota

	// StandardMinus represents the -
	StandardMinus

	// StandardMul represents the *
	StandardMul

	// StandardDiv represents the /
	StandardDiv

	// StandardMod represents the %
	StandardMod

	// StandardSmallerThan represents the <
	StandardSmallerThan

	// StandardSmallerThan represents the <=
	StandardSmallerThanOrEqual

	// StandardBiggerThan represents the >
	StandardBiggerThan

	// StandardBiggerThanOrEqual represents the >=
	StandardBiggerThanOrEqual

	// StandardEqual represents the ==
	StandardEqual

	// StandardNotEqual represents the !=
	StandardNotEqual

	// StandardAnd represents the &&
	StandardAnd

	// StandardOr represents the ||
	StandardOr

	// StandardNot represents the !
	StandardNot

	// StandardXor represents the xor
	StandardXor
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewOperationBuilder creates a new operation builder
func NewOperationBuilder() OperationBuilder {
	hashAdapter := hash.NewAdapter()
	return createOperationBuilder(
		hashAdapter,
	)
}

// NewStandardBuilder creates a new standard builder
func NewStandardBuilder() StandardBuilder {
	hashAdapter := hash.NewAdapter()
	return createStandardBuilder(
		hashAdapter,
	)
}

// NewBitShiftBuilder creates a new bitshift builder
func NewBitShiftBuilder() BitShiftBuilder {
	hashAdapter := hash.NewAdapter()
	return createBitshiftBuilder(
		hashAdapter,
	)
}

// NewSingleSwordBuilder creates a new single sword builder
func NewSingleSwordBuilder() SingleSwordBuilder {
	hashAdapter := hash.NewAdapter()
	return createSingleSwordBuilder(
		hashAdapter,
	)
}

// Builder represents the operations builder
type Builder interface {
	Create() Builder
	WithList(list []Operation) Builder
	Now() (Operations, error)
}

// Operations represents operations
type Operations interface {
	Hash() hash.Hash
	List() []Operation
}

// OperationBuilder represents the operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithStandard(standard Standard) OperationBuilder
	WithSingleSword(singleSword SingleSword) OperationBuilder
	WithBitShift(bitshift BitShift) OperationBuilder
	WithValue(value values.Value) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation
type Operation interface {
	Hash() hash.Hash
	IsStandard() bool
	Standard() Standard
	IsSingleSword() bool
	SingleSword() SingleSword
	IsBitShift() bool
	BitShift() BitShift
	IsValue() bool
	Value() values.Value
}

// StandardBuilder represents the standard operation builder
type StandardBuilder interface {
	Create() StandardBuilder
	WithFirst(first Operation) StandardBuilder
	WithSecond(second Operation) StandardBuilder
	WithFlag(flag uint8) StandardBuilder
	Now() (Standard, error)
}

// Standard represents a standard operation
// arithmetic: +, -, *, /, %
// relational: <, >, <=, >=, ==, !=
// logical: and, or, not, xor
type Standard interface {
	Hash() hash.Hash
	First() Operation
	Second() Operation
	Flag() uint8
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

// SingleSwordBuilder represents the single sword operation builder
type SingleSwordBuilder interface {
	Create() SingleSwordBuilder
	WithVariable(variable string) SingleSwordBuilder
	WithFlag(flag uint8) SingleSwordBuilder
	Now() (SingleSword, error)
}

// SingleSword represents a single sword operation
type SingleSword interface {
	Hash() hash.Hash
	Variable() string
	Flag() uint8 // ++, --
}
