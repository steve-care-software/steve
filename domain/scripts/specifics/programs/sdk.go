package programs

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/functions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/suites"
)

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
