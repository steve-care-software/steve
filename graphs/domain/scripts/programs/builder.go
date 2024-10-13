package programs

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/params"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

type builder struct {
	head         heads.Head
	instructions instructions.Instructions
	params       params.Params
	children     references.References
}

func createBuilder() Builder {
	return &builder{
		head:         nil,
		instructions: nil,
		params:       nil,
		children:     nil,
	}
}

func (app *builder) Create() Builder {
	return createBuilder()
}

func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.instructions = instructions
	return app
}

func (app *builder) WithParams(params params.Params) Builder {
	app.params = params
	return app
}

func (app *builder) WithChildren(children references.References) Builder {
	app.children = children
	return app
}

func (app *builder) Now() (Program, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Program instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions are mandatory in order to build a Program instance")
	}

	if app.params != nil && app.children != nil {
		return createProgramWithParamsAndChildren(
			app.head,
			app.instructions,
			app.params,
			app.children,
		), nil
	}

	if app.params != nil {
		return createProgramWithParams(
			app.head,
			app.instructions,
			app.params,
		), nil
	}

	if app.children != nil {
		return createProgramWithChildren(
			app.head,
			app.instructions,
			app.children,
		), nil
	}

	return createProgram(
		app.head,
		app.instructions,
	), nil
}
