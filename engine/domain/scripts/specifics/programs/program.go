package programs

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/functions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/suites"
	"github.com/steve-care-software/steve/hash"
)

type program struct {
	hash         hash.Hash
	head         heads.Head
	input        string
	instructions instructions.Instructions
	functions    functions.Functions
	suites       suites.Suites
}

func createProgram(
	hash hash.Hash,
	head heads.Head,
	input string,
	instructions instructions.Instructions,
) Program {
	return createProgramInternally(
		hash,
		head,
		input,
		instructions,
		nil,
		nil,
	)
}

func createProgramWithFunctions(
	hash hash.Hash,
	head heads.Head,
	input string,
	instructions instructions.Instructions,
	functions functions.Functions,
) Program {
	return createProgramInternally(
		hash,
		head,
		input,
		instructions,
		functions,
		nil,
	)
}

func createProgramWithSuites(
	hash hash.Hash,
	head heads.Head,
	input string,
	instructions instructions.Instructions,
	suites suites.Suites,
) Program {
	return createProgramInternally(
		hash,
		head,
		input,
		instructions,
		nil,
		suites,
	)
}

func createProgramWithFunctionsAndSuites(
	hash hash.Hash,
	head heads.Head,
	input string,
	instructions instructions.Instructions,
	functions functions.Functions,
	suites suites.Suites,
) Program {
	return createProgramInternally(
		hash,
		head,
		input,
		instructions,
		functions,
		suites,
	)
}

func createProgramInternally(
	hash hash.Hash,
	head heads.Head,
	input string,
	instructions instructions.Instructions,
	functions functions.Functions,
	suites suites.Suites,
) Program {
	out := program{
		hash:         hash,
		head:         head,
		input:        input,
		instructions: instructions,
		functions:    functions,
		suites:       suites,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *program) Head() heads.Head {
	return obj.head
}

// Input returns the head
func (obj *program) Input() string {
	return obj.input
}

// Instructions returns the instructions
func (obj *program) Instructions() instructions.Instructions {
	return obj.instructions
}

// HasFunctions returns true if there is functions, false otherwise
func (obj *program) HasFunctions() bool {
	return obj.functions != nil
}

// Functions returns the functions, if any
func (obj *program) Functions() functions.Functions {
	return obj.functions
}

// HasSuites returns true if there is suites, false otherwise
func (obj *program) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *program) Suites() suites.Suites {
	return obj.suites
}
