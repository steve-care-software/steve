package instructions

import (
	"github.com/steve-care-software/steve/engine/domain/graphs/connections/links"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
	selectors_chain "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

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
	IsRoute() bool
	Route() Route
	IsGrammar() bool
	Grammar() Grammar
	IsAST() bool
	AST() AST
	IsQuery() bool
	Query() AssignableEngineQuery
}

// AssignableEngineQuery represents an assignable engine query
type AssignableEngineQuery interface {
	IsSave() bool
	Save() Save
	IsDelete() bool
	Delete() Delete
	IsSelect() bool
	Select() Select
	IsBridges() bool
	Bridges() Bridges
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

// SingleVariableOperation represents a single variable operation
type SingleVariableOperation interface {
	Name() string
	Oprator() uint8
}

// Casting represents a casting
type Casting interface {
	Variable() string
	Kind() kinds.Kind
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

// NumericValue represents a numeric value
type NumericValue interface {
	IsFloat() bool
	Float() float64
	IsUint() bool
	Uint() *uint64
	IsInt() bool
	Int() *int64
}

// Route represents a route
type Route interface {
	IsOptimal() bool
	Link() links.Link
}

// Grammar represents a grammar
type Grammar interface {
}

// AST represents an ast
type AST interface {
}

// Save represents a save
type Save interface {
}

// Delete represents a delete
type Delete interface {
}

// Select represents a select
type Select interface {
}

// Bridges represents bridges
type Bridges interface {
}
