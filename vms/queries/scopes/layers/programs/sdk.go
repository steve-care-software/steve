package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/programs/reduces"
	"github.com/steve-care-software/steve/vms/queries/scopes/layers/programs/signatures"
)

// Builder represents a programs builder
type Builder interface {
	Create() Builder
	WithList(list []Program) Builder
	Now() (Programs, error)
}

// Programs represents programs
type Programs interface {
	Hash() hash.Hash
	List() []Program
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithInstructions(instructions Instructions) ProgramBuilder
	WithSignature(signature signatures.Signature) ProgramBuilder
	WithSuites(suites Suites) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
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
	WithCompare(compare bytes_programs.Programs) AssignableBuilder
	WithLength(length bytes_programs.Programs) AssignableBuilder
	WithJoin(join bytes_programs.Programs) AssignableBuilder
	WithValue(value bytes_programs.Programs) AssignableBuilder
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
	Compare() bytes_programs.Programs
	IsLength() bool
	Length() bytes_programs.Programs
	IsJoin() bool
	Join() bytes_programs.Programs
	IsValue() bool
	Value() bytes_programs.Programs
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithConstraint(constraint bytes_programs.Program) ConditionBuilder
	WithInstructions(instructions Instructions) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Constraint() bytes_programs.Program
	Instructions() Instructions
}

// QueryBuilder represents the query builder
type QueryBuilder interface {
	Create() QueryBuilder
	WithLayer(layer string) QueryBuilder
	WithBytes(bytes bytes_programs.Program) QueryBuilder
	WithParams(params bytes_programs.Programs) QueryBuilder
	WithDependencies(dependencies []string) QueryBuilder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Layer() string
	Bytes() bytes_programs.Program
	Params() bytes_programs.Programs
	HasDependencies() bool
	Dependencies() []string
}
