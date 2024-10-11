package assignables

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
	selectors_chain "github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/balances/selectors/chains"
)

// NewMapKeyValueBuilder creates a new map key value builder
func NewMapKeyValueBuilder() MapKeyValueBuilder {
	return createMapKeyValueBuilder()
}

// NewAssignableEngineBuilder creates a new assignable engine builderg
func NewAssignableEngineBuilder() AssignableEngineBuilder {
	return createAssignableEngineBuilder()
}

// NewOperationBuilder creates a new operation builder
func NewOperationBuilder() OperationBuilder {
	return createOperationBuilder()
}

// NewOperatorAssignablesBuilder creates a new operator assignablse builder
func NewOperatorAssignablesBuilder() OperatorAssignablesBuilder {
	return createOperatorAssignablesBuilder()
}

// NewOperatorAssignableBuilder creates a new operator assignable builder
func NewOperatorAssignableBuilder() OperatorAssignableBuilder {
	return createOperatorAssignableBuilder()
}

// NewOperatorBuilder creates a new operator builder
func NewOperatorBuilder() OperatorBuilder {
	return createOperatorBuilder()
}

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

// MapKeyValueBuilder represents a map key->value builder
type MapKeyValueBuilder interface {
	Create() MapKeyValueBuilder
	WithName(name string) MapKeyValueBuilder
	WithAssignable(assignable Assignable) MapKeyValueBuilder
	Now() (MapKeyValue, error)
}

// MapKeyValue represents a map key value
type MapKeyValue interface {
	Name() string
	Assignable() Assignable
}

// AssignableEngineBuilder represents an assignable builder
type AssignableEngineBuilder interface {
	Create() AssignableEngineBuilder
	WithSelector(selector selectors_chain.Chain) AssignableEngineBuilder
	WithGrammar(grammar grammars.Grammar) AssignableEngineBuilder
	WithQuery(query queries.Query) AssignableEngineBuilder
	Now() (AssignableEngine, error)
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

// OperationBuilder represents the operation builder
type OperationBuilder interface {
	Create() OperationBuilder
	WithFirst(first Assignable) OperationBuilder
	WithAssignables(assignables OperatorAssignables) OperationBuilder
	Now() (Operation, error)
}

// Operation represents an operation
type Operation interface {
	First() Assignable
	HasAssignables() bool
	Assignables() OperatorAssignables
}

// OperatorAssignablesBuilder represents the operator assignbales builder
type OperatorAssignablesBuilder interface {
	Create() OperatorAssignablesBuilder
	WithList(list []OperatorAssignable) OperatorAssignablesBuilder
	Now() (OperatorAssignables, error)
}

// OperatorAssignables represents an operator assignables
type OperatorAssignables interface {
	List() []OperatorAssignable
}

// OperatorAssignableBuilder represents the operator assignable builder
type OperatorAssignableBuilder interface {
	Create() OperatorAssignableBuilder
	WithOperator(operator Operator) OperatorAssignableBuilder
	WithAssignable(assignable Assignable) OperatorAssignableBuilder
	Now() (OperatorAssignable, error)
}

// OperatorAssignable represents an operator assignable
type OperatorAssignable interface {
	Operator() Operator
	Assignable() Assignable
}

// OperatorBuilder represents the operator builder
type OperatorBuilder interface {
	Create() OperatorBuilder
	WithArithmetic(arithmetic uint8) OperatorBuilder
	WithRelational(relational uint8) OperatorBuilder
	WithEqual(equal uint8) OperatorBuilder
	WithLogical(logical uint8) OperatorBuilder
	Now() (Operator, error)
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
	Logical() *uint8
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
