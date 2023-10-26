package layers

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewLayerBuilder creates a new layer builder
func NewLayerBuilder() LayerBuilder {
	return createLayerBuilder()
}

// NewSuitesBuilder creates a new suites builder
func NewSuitesBuilder() SuitesBuilder {
	return createSuitesBuilder()
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
}

// NewExecutionsBuilder creates a new executions builder
func NewExecutionsBuilder() ExecutionsBuilder {
	return createExecutionsBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	return createAssignmentBuilder()
}

// NewAssignableBuilder create sa new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	return createAssignableBuilder()
}

// NewConstantValuesBuilder creates a new constant values builder
func NewConstantValuesBuilder() ConstantValuesBuilder {
	return createConstantValuesBuilder()
}

// NewConstantValueBuilder creates a new constant value builder
func NewConstantValueBuilder() ConstantValueBuilder {
	return createConstantValueBuilder()
}

// NewConditionBuilder creates a new condition instance
func NewConditionBuilder() ConditionBuilder {
	return createConditionBuilder()
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() QueryBuilder {
	return createQueryBuilder()
}

// NewLayerInputBuilder creates a new layer input builder
func NewLayerInputBuilder() LayerInputBuilder {
	return createLayerInputBuilder()
}

// NewValueAssignmentsBuilder creates a new value assignments builder
func NewValueAssignmentsBuilder() ValueAssignmentsBuilder {
	return createValueAssignmentsBuilder()
}

// NewValueAssignmentBuilder creates a new value assignment builder
func NewValueAssignmentBuilder() ValueAssignmentBuilder {
	return createValueAssignmentBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents a layers builder
type Builder interface {
	Create() Builder
	WithList(list []Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	List() []Layer
}

// LayerBuilder represents a layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithInput(input string) LayerBuilder
	WithExecutions(executions Executions) LayerBuilder
	WithReturn(ret return_expectations.Expectation) LayerBuilder
	WithParams(params parameters.Parameters) LayerBuilder
	WithSuites(suites Suites) LayerBuilder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Input() string
	Executions() Executions
	Return() return_expectations.Expectation
	HasParams() bool
	Params() parameters.Parameters
	HasSuites() bool
	Suites() Suites
}

// SuitesBuilder represents a suites builder
type SuitesBuilder interface {
	Create() SuitesBuilder
	WithList(list []Suite) SuitesBuilder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithInput(input []byte) SuiteBuilder
	WithReturn(ret returns.Return) SuiteBuilder
	WithValues(values ValueAssignments) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a layer suites
type Suite interface {
	Name() string
	Input() []byte
	Return() returns.Return
	HasValues() bool
	Values() ValueAssignments
}

// ExecutionsBuilder represents an executions builder
type ExecutionsBuilder interface {
	Create() ExecutionsBuilder
	WithList(list []Execution) ExecutionsBuilder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithAssignment(assignmnet Assignment) ExecutionBuilder
	WithCondition(condition Condition) ExecutionBuilder
	IsStop() ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	IsStop() bool
	IsAssignment() bool
	Assignment() Assignment
	IsCondition() bool
	Condition() Condition
}

// AssignmentBuilder represents an assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithQuery(query Query) AssignableBuilder
	WithReduce(reduce reduces.Reduce) AssignableBuilder
	WithCompare(compare ConstantValues) AssignableBuilder
	WithLength(length ConstantValue) AssignableBuilder
	WithJoin(join ConstantValues) AssignableBuilder
	WithValue(value ConstantValue) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsQuery() bool
	Query() Query
	IsReduce() bool
	Reduce() reduces.Reduce
	IsCompare() bool
	Compare() ConstantValues
	IsLength() bool
	Length() ConstantValue
	IsJoin() bool
	Join() ConstantValues
	IsValue() bool
	Value() ConstantValue
}

// ConstantValuesBuilder represents constant values builder
type ConstantValuesBuilder interface {
	Create() ConstantValuesBuilder
	WithList(list []ConstantValue) ConstantValuesBuilder
	Now() (ConstantValues, error)
}

// ConstantValues represents constant values
type ConstantValues interface {
	List() []ConstantValue
}

// ConstantValueBuilder represents a constant value builder
type ConstantValueBuilder interface {
	Create() ConstantValueBuilder
	WithVariable(variable string) ConstantValueBuilder
	WithConstant(constant []byte) ConstantValueBuilder
	Now() (ConstantValue, error)
}

// ConstantValue represents a constant value
type ConstantValue interface {
	IsVariable() bool
	Variable() string
	IsConstant() bool
	Constant() []byte
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithVariable(variable string) ConditionBuilder
	WithExecutions(executions Executions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Variable() string
	Executions() Executions
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithInput(input ConstantValue) QueryBuilder
	WithLayer(layer LayerInput) QueryBuilder
	WithValues(values ValueAssignments) QueryBuilder
	Now() (Query, error)
}

// Query represents a query execution
type Query interface {
	Input() ConstantValue
	Layer() LayerInput
	HasValues() bool
	Values() ValueAssignments
}

// LayerInputBuilder represents a layer input builder
type LayerInputBuilder interface {
	Create() LayerInputBuilder
	WithVariable(variable string) LayerInputBuilder
	WithLayer(layer Layer) LayerInputBuilder
	Now() (LayerInput, error)
}

// LayerInput represents a layer input
type LayerInput interface {
	IsVariable() bool
	Variable() string
	IsLayer() bool
	Layer() Layer
}

// ValueAssignmentsBuilder represents a value assignmnets builder
type ValueAssignmentsBuilder interface {
	Create() ValueAssignmentsBuilder
	WithList(list []ValueAssignment) ValueAssignmentsBuilder
	Now() (ValueAssignments, error)
}

// ValueAssignments represents a value assignments
type ValueAssignments interface {
	List() []ValueAssignment
}

// ValueAssignmentBuilder repreents a value assignment builder
type ValueAssignmentBuilder interface {
	Create() ValueAssignmentBuilder
	WithName(name string) ValueAssignmentBuilder
	WithValue(value Value) ValueAssignmentBuilder
	Now() (ValueAssignment, error)
}

// ValueAssignment represents a alue assignment
type ValueAssignment interface {
	Name() string
	Value() Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithVariable(variable string) ValueBuilder
	WithConstant(constant []byte) ValueBuilder
	WithLayer(layer Layer) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsVariable() bool
	Variable() string
	IsConstant() bool
	Constant() []byte
	IsLayer() bool
	Layer() Layer
}
