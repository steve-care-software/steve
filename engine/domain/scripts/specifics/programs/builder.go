package programs

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/functions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/suites"
)

type builder struct {
	hashAdapter  hash.Adapter
	head         heads.Head
	input        string
	instructions instructions.Instructions
	functions    functions.Functions
	suites       suites.Suites
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		head:         nil,
		input:        "",
		instructions: nil,
		functions:    nil,
		suites:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// WithInstructions adds an instructions to the builder
func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.instructions = instructions
	return app
}

// WithFunctions adds an functions to the builder
func (app *builder) WithFunctions(functions functions.Functions) Builder {
	app.functions = functions
	return app
}

// WithSuites adds a suites to the builder
func (app *builder) WithSuites(suites suites.Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Program instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Program instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Program instance")
	}

	data := [][]byte{
		app.head.Hash().Bytes(),
		[]byte(app.input),
		app.instructions.Hash().Bytes(),
	}

	if app.functions != nil {
		data = append(data, app.functions.Hash().Bytes())
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.functions != nil && app.suites != nil {
		return createProgramWithFunctionsAndSuites(
			*pHash,
			app.head,
			app.input,
			app.instructions,
			app.functions,
			app.suites,
		), nil
	}

	if app.functions != nil {
		return createProgramWithFunctions(
			*pHash,
			app.head,
			app.input,
			app.instructions,
			app.functions,
		), nil
	}

	if app.suites != nil {
		return createProgramWithSuites(
			*pHash,
			app.head,
			app.input,
			app.instructions,
			app.suites,
		), nil
	}

	return createProgram(
		*pHash,
		app.head,
		app.input,
		app.instructions,
	), nil
}
