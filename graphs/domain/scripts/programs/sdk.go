package programs

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/params"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithHead(head heads.Head) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithParams(params params.Params) Builder
	WithChildren(children references.References) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Head() heads.Head
	Instructions() instructions.Instructions
	HasParams() bool
	Params() params.Params
	HasChildren() bool
	Children() references.References
}
