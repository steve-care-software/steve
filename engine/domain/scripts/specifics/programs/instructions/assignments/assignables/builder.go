package assignables

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
)

type builder struct {
	hashAdapter hash.Adapter
	operation   operations.Operation
	call        calls.Call
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		operation:   nil,
		call:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOperation adds an operation to the builder
func (app *builder) WithOperation(operation operations.Operation) Builder {
	app.operation = operation
	return app
}

// WithCall adds a call to the builder
func (app *builder) WithCall(call calls.Call) Builder {
	app.call = call
	return app
}

// Now builds a new Assignable instance
func (app *builder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.operation != nil {
		data = append(data, app.operation.Hash().Bytes())
	}

	if app.call != nil {
		data = append(data, app.call.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Assignable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.operation != nil {
		return createAssignableWithOperation(*pHash, app.operation), nil
	}

	return createAssignableWithCall(*pHash, app.call), nil
}
