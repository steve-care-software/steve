package functions

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/functions/parameters"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/suites"
)

// Functions represents functions
type Functions interface {
	Hash() hash.Hash
	List() []Function
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
