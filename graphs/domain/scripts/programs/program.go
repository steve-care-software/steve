package programs

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/params"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

type program struct {
	head         heads.Head
	instructions instructions.Instructions
	params       params.Params
	children     references.References
}

func createProgram(
	head heads.Head,
	instructions instructions.Instructions,
) Program {
	return createProgramInternally(
		head,
		instructions,
		nil,
		nil,
	)
}

func createProgramWithParams(
	head heads.Head,
	instructions instructions.Instructions,
	params params.Params,
) Program {
	return createProgramInternally(
		head,
		instructions,
		params,
		nil,
	)
}

func createProgramWithChildren(
	head heads.Head,
	instructions instructions.Instructions,
	children references.References,
) Program {
	return createProgramInternally(
		head,
		instructions,
		nil,
		children,
	)
}

func createProgramWithParamsAndChildren(
	head heads.Head,
	instructions instructions.Instructions,
	params params.Params,
	children references.References,
) Program {
	return createProgramInternally(
		head,
		instructions,
		params,
		children,
	)
}

func createProgramInternally(
	head heads.Head,
	instructions instructions.Instructions,
	params params.Params,
	children references.References,
) Program {
	out := program{
		head:         head,
		instructions: instructions,
		params:       params,
		children:     children,
	}

	return &out
}

func (obj *program) Head() heads.Head {
	return obj.head
}

func (obj *program) Instructions() instructions.Instructions {
	return obj.instructions
}

func (obj *program) HasParams() bool {
	return obj.params != nil
}

func (obj *program) Params() params.Params {
	return obj.params
}

func (obj *program) HasChildren() bool {
	return obj.children != nil
}

func (obj *program) Children() references.References {
	return obj.children
}
