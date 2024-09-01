package programs

import (
	"github.com/steve-care-software/steve/domain/programs/instructions"
)

type program struct {
	root instructions.Element
}

func createProgram(
	root instructions.Element,
) Program {
	return createProgramInternally(root)
}

func createProgramInternally(
	root instructions.Element,
) Program {
	out := program{
		root: root,
	}

	return &out
}

// Root returns the root
func (obj *program) Root() instructions.Element {
	return obj.root
}
