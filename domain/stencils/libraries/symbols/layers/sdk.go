package layers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns"
	return_expectations "github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/expectations"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLayerBuilder creates a new layer builder
func NewLayerBuilder() LayerBuilder {
	hashAdapter := hash.NewAdapter()
	return createLayerBuilder(
		hashAdapter,
	)
}

// NewSuitesBuilder creates a new suites builder
func NewSuitesBuilder() SuitesBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuitesBuilder(
		hashAdapter,
	)
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	hashAdapter := hash.NewAdapter()
	return createSuiteBuilder(
		hashAdapter,
	)
}

// NewExecutionsBuilder creates a new executions builder
func NewExecutionsBuilder() ExecutionsBuilder {
	hashAdapter := hash.NewAdapter()
	return createExecutionsBuilder(
		hashAdapter,
	)
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	hashAdapter := hash.NewAdapter()
	return createExecutionBuilder(
		hashAdapter,
	)
}

// NewAssignmentBuilder creates a new assignment builder
func NewAssignmentBuilder() AssignmentBuilder {
	hashAdapter := hash.NewAdapter()
	return createAssignmentBuilder(
		hashAdapter,
	)
}

// NewAssignableBuilder create sa new assignable builder
func NewAssignableBuilder() AssignableBuilder {
	return createAssignableBuilder()
}

// NewConditionBuilder creates a new condition instance
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() QueryBuilder {
	hashAdapter := hash.NewAdapter()
	return createQueryBuilder(
		hashAdapter,
	)
}

// NewLayerInputBuilder creates a new layer input builder
func NewLayerInputBuilder() LayerInputBuilder {
	hashAdapter := hash.NewAdapter()
	return createLayerInputBuilder(
		hashAdapter,
	)
}

// NewValueAssignmentsBuilder creates a new value assignments builder
func NewValueAssignmentsBuilder() ValueAssignmentsBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueAssignmentsBuilder(
		hashAdapter,
	)
}

// NewValueAssignmentBuilder creates a new value assignment builder
func NewValueAssignmentBuilder() ValueAssignmentBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueAssignmentBuilder(
		hashAdapter,
	)
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueBuilder(
		hashAdapter,
	)
}

// Builder represents a layers builder
type Builder interface {
	Create() Builder
	WithList(list []Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents an assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithQuery(query Query) AssignableBuilder
	WithReduce(reduce reduces.Reduce) AssignableBuilder
	WithCompare(compare constantvalues.ConstantValues) AssignableBuilder
	WithLength(length constantvalues.ConstantValue) AssignableBuilder
	WithJoin(join constantvalues.ConstantValues) AssignableBuilder
	WithValue(value constantvalues.ConstantValue) AssignableBuilder
	WithAccount(account accounts.Account) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsQuery() bool
	Query() Query
	IsReduce() bool
	Reduce() reduces.Reduce
	IsCompare() bool
	Compare() constantvalues.ConstantValues
	IsLength() bool
	Length() constantvalues.ConstantValue
	IsJoin() bool
	Join() constantvalues.ConstantValues
	IsValue() bool
	Value() constantvalues.ConstantValue
	IsAccount() bool
	Account() accounts.Account
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
	Hash() hash.Hash
	Variable() string
	Executions() Executions
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithInput(input constantvalues.ConstantValue) QueryBuilder
	WithLayer(layer LayerInput) QueryBuilder
	WithValues(values ValueAssignments) QueryBuilder
	Now() (Query, error)
}

// Query represents a query execution
type Query interface {
	Hash() hash.Hash
	Input() constantvalues.ConstantValue
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
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
	Hash() hash.Hash
	IsVariable() bool
	Variable() string
	IsConstant() bool
	Constant() []byte
	IsLayer() bool
	Layer() Layer
}
