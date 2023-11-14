package links

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/pointers"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers"
)

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithOrigin(origin Origin) Builder
	WithInstructions(instructions Instructions) Builder
	WithExecution(execution Query) Builder
	WithSuites(suites Suites) Builder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Origin() Origin
	Instructions() Instructions
	Execution() layers.Query
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
	WithQuery(query Query) SuiteBuilder
	WithReturns(returns []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Name() string
	Query() Query
	Returns() []byte
}

// QueryBuilder represents a query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithLayer(layer LayerValue) QueryBuilder
	WithBytes(bytes layers.ByteValue) QueryBuilder
	WithParams(params layers.ByteValues) QueryBuilder
	WithDependencies(dependencies LayerValues) QueryBuilder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Layer() LayerValue
	Bytes() layers.ByteValue
	Params() layers.ByteValues
	HasDependencies() bool
	Dependencies() LayerValues
}

// LayerValuesBuilder represents the layer values builder
type LayerValuesBuilder interface {
	Create() LayerValuesBuilder
	WithList(list []LayerValue) LayerValuesBuilder
	Now() (LayerValues, error)
}

// LayerValues represents layer values
type LayerValues interface {
	Hash() hash.Hash
	List() []LayerValue
}

// LayerValueBuilder represents a layer value builder
type LayerValueBuilder interface {
	Create() LayerValueBuilder
	WithVariable(variable string) LayerValueBuilder
	WithPointer(pointer pointers.Pointer) LayerValueBuilder
	Now() (LayerValue, error)
}

// LayerValue represents a layer value
type LayerValue interface {
	Hash() hash.Hash
	IsVariable() bool
	Variable() string
	IsPointer() bool
	Pointer() pointers.Pointer
}

// InstructionsBuilder represents the instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	Hash() hash.Hash
	List() []Instruction
}

// InstructionBuilder represents an instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment InstructionBuilder) InstructionBuilder
	WithCondition(condition Condition) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Hash() hash.Hash
	IsAssignment() bool
	Assignment() Assignment
	IsCondition() bool
	Condition() Condition
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithConstraint(constraint layers.ByteValue) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Constraint() layers.ByteValue
	Instructions() Instructions
}

// AssignmentBuilder represents the assignment builder
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

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithLoad(load pointers.Pointer) AssignableBuilder
	WithExists(exists pointers.Pointer) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	Hash() hash.Hash
	IsLoad() bool
	Load() pointers.Pointer
	IsExists() bool
	Exists() pointers.Pointer
}

// OriginBuilder represents the origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithPointer(pointer pointers.Pointer) OriginBuilder
	WithNext(next Origin) OriginBuilder
	Now() (Origin, error)
}

// Origin represents a link origin
type Origin interface {
	Hash() hash.Hash
	Pointer() pointers.Pointer
	HasNext() bool
	Next() Origin
}
