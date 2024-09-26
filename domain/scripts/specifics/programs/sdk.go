package programs

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/suites"
)

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Head() heads.Head
	Input() string
	Instructions() instructions.Instructions
	HasSuites() bool
	Suites() suites.Suites
}

// Function represents a function
type Function interface {
	Hash() hash.Hash
	Parameters() FuncParameters
	Instructions() instructions.Instructions
	HasOutput() bool
	Output() containers.Containers
	HasSuites() bool
	Suites() suites.Suites
}

// FuncParameters represents func parameters
type FuncParameters interface {
	Hash() hash.Hash
	List() []FuncParameter
}

// FuncParameter represents a func parameter
type FuncParameter interface {
	Hash() hash.Hash
	Name() string
	Container() containers.Container
	IsMandatory() bool
}
