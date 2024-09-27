package functions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/functions/parameters"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/suites"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewFunctionBuilder creates a new function builder
func NewFunctionBuilder() FunctionBuilder {
	hashAdapter := hash.NewAdapter()
	return createFunctionBuilder(
		hashAdapter,
	)
}

// Builder represents a functions builder
type Builder interface {
	Create() Builder
	WithList(list []Function) Builder
	Now() (Functions, error)
}

// Functions represents functions
type Functions interface {
	Hash() hash.Hash
	List() []Function
}

// FunctionBuilder represents the function builder
type FunctionBuilder interface {
	Create() FunctionBuilder
	WithParameters(parameters parameters.Parameters) FunctionBuilder
	WithInstructions(instructions instructions.Instructions) FunctionBuilder
	WithOutput(output containers.Containers) FunctionBuilder
	WithSuites(suites suites.Suites) FunctionBuilder
	Now() (Function, error)
}

// Function represents a function
type Function interface {
	Hash() hash.Hash
	Parameters() parameters.Parameters
	Instructions() instructions.Instructions
	HasOutput() bool
	Output() containers.Containers
	HasSuites() bool
	Suites() suites.Suites
}
