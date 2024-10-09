package programs

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/params"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

// Program represents a program
type Program interface {
	Head() heads.Head
	Instructions() instructions.Instructions
	HasParams() bool
	Params() params.Params
	HasChildren() bool
	Children() references.References
}
