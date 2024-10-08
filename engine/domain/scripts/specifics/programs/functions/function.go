package functions

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/functions/parameters"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/suites"
	"github.com/steve-care-software/steve/hash"
)

type function struct {
	hash         hash.Hash
	parameters   parameters.Parameters
	instructions instructions.Instructions
	output       containers.Containers
	suites       suites.Suites
}

func createFunction(
	hash hash.Hash,
	parameters parameters.Parameters,
	instructions instructions.Instructions,
) Function {
	return createFunctionInternally(
		hash,
		parameters,
		instructions,
		nil,
		nil,
	)
}

func createFunctionWithOutput(
	hash hash.Hash,
	parameters parameters.Parameters,
	instructions instructions.Instructions,
	output containers.Containers,
) Function {
	return createFunctionInternally(
		hash,
		parameters,
		instructions,
		output,
		nil,
	)
}

func createFunctionWithSuites(
	hash hash.Hash,
	parameters parameters.Parameters,
	instructions instructions.Instructions,
	suites suites.Suites,
) Function {
	return createFunctionInternally(
		hash,
		parameters,
		instructions,
		nil,
		suites,
	)
}

func createFunctionWithOutputAndSuites(
	hash hash.Hash,
	parameters parameters.Parameters,
	instructions instructions.Instructions,
	output containers.Containers,
	suites suites.Suites,
) Function {
	return createFunctionInternally(
		hash,
		parameters,
		instructions,
		output,
		suites,
	)
}

func createFunctionInternally(
	hash hash.Hash,
	parameters parameters.Parameters,
	instructions instructions.Instructions,
	output containers.Containers,
	suites suites.Suites,
) Function {
	out := function{
		hash:         hash,
		parameters:   parameters,
		instructions: instructions,
		output:       output,
		suites:       suites,
	}

	return &out
}

// Hash returns the hash
func (obj *function) Hash() hash.Hash {
	return obj.hash
}

// Parameters returns the parameters
func (obj *function) Parameters() parameters.Parameters {
	return obj.parameters
}

// Instructions returns the instructions
func (obj *function) Instructions() instructions.Instructions {
	return obj.instructions
}

// HasOutput returns true if there is an output, false otherwise
func (obj *function) HasOutput() bool {
	return obj.output != nil
}

// Output returns the output, if any
func (obj *function) Output() containers.Containers {
	return obj.output
}

// HasSuites returns true if there is suites, false otherwise
func (obj *function) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *function) Suites() suites.Suites {
	return obj.suites
}
