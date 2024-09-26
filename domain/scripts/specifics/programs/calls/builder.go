package calls

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/engines"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/functions"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/calls/programs"
)

type builder struct {
	hashAdapter hash.Adapter
	program     programs.Program
	engine      engines.Engine
	function    functions.Function
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		program:     nil,
		engine:      nil,
		function:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program programs.Program) Builder {
	app.program = program
	return app
}

// WithEngine adds an engine to the builder
func (app *builder) WithEngine(engine engines.Engine) Builder {
	app.engine = engine
	return app
}

// WithFunction adds a function to the builder
func (app *builder) WithFunction(function functions.Function) Builder {
	app.function = function
	return app
}

// Now builds a new Call instance
func (app *builder) Now() (Call, error) {
	data := [][]byte{}
	if app.program != nil {
		data = append(data, app.program.Hash().Bytes())
	}

	if app.engine != nil {
		data = append(data, app.engine.Hash().Bytes())
	}

	if app.function != nil {
		data = append(data, app.function.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Call is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.program != nil {
		return createCallWithProgram(
			*pHash,
			app.program,
		), nil
	}

	if app.engine != nil {
		return createCallWithEngine(
			*pHash,
			app.engine,
		), nil
	}

	return createCallWithFunction(
		*pHash,
		app.function,
	), nil

}
