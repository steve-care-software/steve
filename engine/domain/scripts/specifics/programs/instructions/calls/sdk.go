package calls

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/engines"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/functions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/programs"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the call builder
type Builder interface {
	Create() Builder
	WithProgram(program programs.Program) Builder
	WithEngine(engine engines.Engine) Builder
	WithFunction(function functions.Function) Builder
	Now() (Call, error)
}

// Call represents a call
type Call interface {
	Hash() hash.Hash
	IsProgram() bool
	Program() programs.Program
	IsEngine() bool
	Engine() engines.Engine
	IsFunction() bool
	Function() functions.Function
}
