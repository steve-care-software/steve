package instructions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
	selectors_chain "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

// NewSingleVariableOperationBuilder creates a new single variable operation builder
func NewSingleVariableOperationBuilder() SingleVariableOperationBuilder {
	return createSingleVariableOperationBuilder()
}

// NewCastingBuilder creates a new casting builder
func NewCastingBuilder() CastingBuilder {
	return createCastingBuilder()
}

// NewPrimitiveValueBuilder creates a new primitive value builderg
func NewPrimitiveValueBuilder() PrimitiveValueBuilder {
	return createPrimitiveValueBuilder()
}

// NewNumericValueBuilder creates a new numeric value builder
func NewNumericValueBuilder() NumericValueBuilder {
	return createNumericValueBuilder()
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// Instruction represents an instruction
type Instruction interface {
	IsSingleVariableOperation() bool
	SingleVariableOperation() SingleVariableOperation
	IsAssignment() bool
	Assignment() Assignment
	IsCondition() bool
	Condition() Condition
	IsProgramCall() bool
	ProgramCall() ProgramCall
	IsForLoop() bool
	ForLoop() ForLoop
	IsReturnInstruction() bool
	ReturnInstruction() ReturnInstruction
}

// ReturnInstruction represents a return instruction
type ReturnInstruction interface {
	HasAssignable() bool
	Assignable() Assignable
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
	ListMap() ListMap
	IsVariable() bool
	Variable() string
}

// ProgramCall represents a program call
type ProgramCall interface {
	References() references.References
	HasParams() bool
	Params() MapKeyValues
}

// ListMap represents a list map
type ListMap interface {
	IsList() bool
	List() Assignables
	IsMap() bool
	Map() MapKeyValues
}

// MapKeyValues represents a map key->values
type MapKeyValues interface {
	List() []MapKeyValue
}

// MapKeyValue represents a map key value
type MapKeyValue interface {
	Name() string
	Assignable() Assignable
}

// Condition represents a condition
type Condition interface {
	Operation() Operation
	Instructions() Instructions
}

// Assignment represents an assignment
type Assignment interface {
	IsMultiple() bool
	Multiple() AssignmentMultiple
	IsOperation() bool
	Operation() AssignmentOperation
}

// AssignmentMultiple represents a multiple assignment
type AssignmentMultiple interface {
	Assignees() Assignees
	Assignables() Assignables
}

// AssignmentOperation represents an assignment operation (ex: myVAr++)
type AssignmentOperation interface {
	Assignee() Assignee
	Operator() uint8        // arithmetic operator
	Assignable() Assignable // must compute to an arithmetic-valid value
}

// Assignees represents assignees
type Assignees interface {
	List() []Assignee
}

// Assignee represents an assignee
type Assignee interface {
	Name() AssigneeName
	HasKind() bool
	Kind() kinds.Kind
}

// AssigneeName represents an assignee name
type AssigneeName interface {
	Name() string
	HasReferences() bool
	References() references.References
	HasIndex() bool
	Index() *uint
}

// Assignables represents assignables
type Assignables interface {
	List() []Assignable
}

// Assignable represents an assignable
type Assignable interface {
	IsEngine() bool
	Engine() AssignableEngine
	IsListMap() bool
	ListMap() ListMap
	IsProgramCall() bool
	ProgramCall() ProgramCall
	IsPrimitive() bool
	Primitive() PrimitiveValue
	IsCasting() bool
	Casting() Casting
	IsExpand() bool
	Expand() string
	IsOperation() bool
	Operation() Operation
}

// AssignableEngine represents an engine assignable
type AssignableEngine interface {
	IsSelector() bool
	Selector() selectors_chain.Chain
	IsGrammar() bool
	Grammar() grammars.Grammar
	IsQuery() bool
	Query() queries.Query
}

// Operation represents an operation
type Operation interface {
	First() Assignable
	HasAssignables() bool
	Assignables() OperatorAssignables
}

// OperatorAssignables represents an operator assignables
type OperatorAssignables interface {
	List() []OperatorAssignable
}

// OperatorAssignable represents an operator assignable
type OperatorAssignable interface {
	Operator() Operator
	Assignable() Assignable
}

// Operator represents an operator
type Operator interface {
	IsArithmetic() bool
	Arithmetic() *uint8
	IsRelational() bool
	Relational() *uint8
	IsEqual() bool
	Equal() *uint8
	IsLogical() bool
	Logical() uint8
}

// SingleVariableOperationBuilder represents a single variable operation builder
type SingleVariableOperationBuilder interface {
	Create() SingleVariableOperationBuilder
	WithName(name string) SingleVariableOperationBuilder
	WithOperator(operator uint8) SingleVariableOperationBuilder
	Now() (SingleVariableOperation, error)
}

// SingleVariableOperation represents a single variable operation
type SingleVariableOperation interface {
	Name() string
	Operator() uint8
}

// CastingBuilder repreents a casting builder
type CastingBuilder interface {
	Create() CastingBuilder
	WithAssignable(assignable assignables.Assignable) CastingBuilder
	WithKind(kind kinds.Kind) CastingBuilder
	Now() (Casting, error)
}

// Casting represents a casting
type Casting interface {
	Assignable() assignables.Assignable
	Kind() kinds.Kind
}

// PrimitiveValueBuilder represents a primitive value builder
type PrimitiveValueBuilder interface {
	Create() PrimitiveValueBuilder
	WithNumeric(numeric NumericValue) PrimitiveValueBuilder
	WithBool(boolBalue bool) PrimitiveValueBuilder
	WithString(strValue string) PrimitiveValueBuilder
	Now() (PrimitiveValue, error)
}

// PrimitiveValue represents a primitive value
type PrimitiveValue interface {
	IsNumeric() bool
	Numeric() NumericValue
	IsBool() bool
	Bool() *bool
	IsString() bool
	String() *string
}

// NumericValueBuilder represents the numeric value builder
type NumericValueBuilder interface {
	Create() NumericValueBuilder
	WithFloat(flValue float64) NumericValueBuilder
	WithUint(uiValue uint64) NumericValueBuilder
	WithInt(intValue int64) NumericValueBuilder
	Now() (NumericValue, error)
}

// NumericValue represents a numeric value
type NumericValue interface {
	IsFloat() bool
	Float() float64
	IsUint() bool
	Uint() *uint64
	IsInt() bool
	Int() *int64
}
