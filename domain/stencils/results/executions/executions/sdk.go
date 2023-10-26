package executions

import "github.com/steve-care-software/steve/domain/stencils/results/executions/executions/assignments"

// Builder represents an executions builder
type Builder interface {
	Create() Builder
	WithList(list []Execution) Builder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
}

// ExecutionBuilder represents the execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithIndex(index uint) ExecutionBuilder
	WithAssignment(assignment assignments.Assignment) ExecutionBuilder
	WithExecutions(executions Executions) ExecutionBuilder
	IsStop() ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Index() uint
	Content() Content
}

// Content represents an execution content
type Content interface {
	IsStop() bool
	IsAssignment() bool
	Assignment() assignments.Assignment
	IsExecutions() bool
	Executions() Executions
}
