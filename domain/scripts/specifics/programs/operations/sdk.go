package operations

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/values"
)

// Operations represents operations
type Operations interface {
	List() Operation
}

// Operation represents an operation
// arithmetic: +, -, *, /, %
// relational: <, >, <=, >=, ==, !=
// logical: and, or, not, xor
type Operation interface {
	Hash() hash.Hash
	IsArithmetic() bool
	Arithmetic() StandardOperation
	IsRelational() bool
	Relational() StandardOperation
	IsLogical() bool
	Logical() StandardOperation
	IsSingleSword() bool
	SingleSword() SingleSwordOperation
	IsBitShift() bool
	BitShift() BitShift
	IsValue() bool
	Value() values.Value
}

// BitShift represents a bitshift
type BitShift interface {
	Operation() Operation
	Flag() uint8 // <<, >>
	Amount() uint8
}

// StandardOperation represents a standard operation
type StandardOperation interface {
	Operations() Operations
	Flag() uint8
}

// SingleSwordOperation represents a single sword operation
type SingleSwordOperation interface {
	Variable() string
	Flag() uint8 // ++, --
}
