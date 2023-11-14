package layers

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers/reduces"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/symbols/layers/signatures"
)

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
	WithInstructions(instructions Instructions) LayerBuilder
	WithSignature(signature signatures.Signature) LayerBuilder
	WithSuites(suites Suites) LayerBuilder
	Now() (Layer, error)
}

// Layer represents a stencil layer
type Layer interface {
	Hash() hash.Hash
	Instructions() Instructions
	Signature() signatures.Signature
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
	WithValid(valid []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Name() string
	Query() Query
	IsValid() bool
	Valid() []byte
}

// InstructionsBuilder represents an instructions builder
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

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	WithCondition(condition Condition) InstructionBuilder
	IsStop() InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
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

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithQuery(query Query) AssignableBuilder
	WithReduce(reduce reduces.Reduce) AssignableBuilder
	WithCompare(compare ByteValues) AssignableBuilder
	WithLength(length ByteValues) AssignableBuilder
	WithJoin(join ByteValues) AssignableBuilder
	WithValue(value ByteValues) AssignableBuilder
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
	Compare() ByteValues
	IsLength() bool
	Length() ByteValues
	IsJoin() bool
	Join() ByteValues
	IsValue() bool
	Value() ByteValues
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithConstraint(constraint ByteValue) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Constraint() ByteValue
	Instructions() Instructions
}

// QueryBuilder represents the query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithLayer(layer LayerValue) QueryBuilder
	WithBytes(bytes ByteValue) QueryBuilder
	WithParams(params ByteValues) QueryBuilder
	WithDependencies(dependencies LayerValues) QueryBuilder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Layer() LayerValue
	Bytes() ByteValue
	Params() ByteValues
	HasDependencies() bool
	Dependencies() LayerValues
}

// ByteValuesBuilder represents a byte valyes builder
type ByteValuesBuilder interface {
	Create() ByteValuesBuilder
	WithList(list []ByteValue) ByteValuesBuilder
	Now() (ByteValues, error)
}

// ByteValues represents byte values
type ByteValues interface {
	Hash() hash.Hash
	List() []ByteValue
}

// ByteValueBuilder represents a byte value builder
type ByteValueBuilder interface {
	Create() ByteValueBuilder
	WithVariable(variable string) ByteValueBuilder
	WithBytes(bytes []byte) ByteValueBuilder
	Now() (ByteValue, error)
}

// ByteValue represents a byte value
type ByteValue interface {
	Hash() hash.Hash
	IsVariable() bool
	Variable() string
	IsBytes() bool
	Bytes() []byte
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

// LayerValueBuilder represents the layer value builder
type LayerValueBuilder interface {
	Create() LayerValueBuilder
	WithVariable(variable string) LayerValueBuilder
	WithLayer(layer Layer) LayerValueBuilder
	Now() (LayerValue, error)
}

// LayerValue represents a layer value
type LayerValue interface {
	Hash() hash.Hash
	IsVariable() bool
	Variable() string
	IsLayer() bool
	Layer() Layer
}
