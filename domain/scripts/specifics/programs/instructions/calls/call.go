package calls

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/calls/engines"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/calls/functions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/calls/programs"
)

type call struct {
	hash     hash.Hash
	program  programs.Program
	engine   engines.Engine
	function functions.Function
}

func createCallWithProgram(
	hash hash.Hash,
	program programs.Program,
) Call {
	return createCallInternally(hash, program, nil, nil)
}

func createCallWithEngine(
	hash hash.Hash,
	engine engines.Engine,
) Call {
	return createCallInternally(hash, nil, engine, nil)
}

func createCallWithFunction(
	hash hash.Hash,
	function functions.Function,
) Call {
	return createCallInternally(hash, nil, nil, function)
}

func createCallInternally(
	hash hash.Hash,
	program programs.Program,
	engine engines.Engine,
	function functions.Function,
) Call {
	out := call{
		hash:     hash,
		program:  program,
		engine:   engine,
		function: function,
	}

	return &out
}

// Hash returns the hash
func (obj *call) Hash() hash.Hash {
	return obj.hash
}

// IsProgram true if there is a program, false otherwise
func (obj *call) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *call) Program() programs.Program {
	return obj.program
}

// IsEngine true if there is an engine, false otherwise
func (obj *call) IsEngine() bool {
	return obj.engine != nil
}

// Engine returns the engine, if any
func (obj *call) Engine() engines.Engine {
	return obj.engine
}

// IsFunction true if there is a function, false otherwise
func (obj *call) IsFunction() bool {
	return obj.function != nil
}

// Function returns the function, if any
func (obj *call) Function() functions.Function {
	return obj.function
}
