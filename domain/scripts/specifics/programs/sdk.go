package programs

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/functions"
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

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithInput(input string) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithFunctions(functions functions.Functions) Builder
	WithSuites(suites suites.Suites) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Head() heads.Head
	Input() string
	Instructions() instructions.Instructions
	HasFunctions() bool
	Functions() functions.Functions
	HasSuites() bool
	Suites() suites.Suites
}
