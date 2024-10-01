package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
)

type loopCounterBuilder struct {
	hashAdapter hash.Adapter
	assignment  assignments.Assignment
	operation   operations.Operation
	increment   operations.Operation
}

func createLoopCounterBuilder(
	hashAdapter hash.Adapter,
) LoopCounterBuilder {
	out := loopCounterBuilder{
		hashAdapter: hashAdapter,
		assignment:  nil,
		operation:   nil,
		increment:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *loopCounterBuilder) Create() LoopCounterBuilder {
	return createLoopCounterBuilder(
		app.hashAdapter,
	)
}

// WithAssignment adds an assignment to the builder
func (app *loopCounterBuilder) WithAssignment(assignment assignments.Assignment) LoopCounterBuilder {
	app.assignment = assignment
	return app
}

// WithOperation adds an operation to the builder
func (app *loopCounterBuilder) WithOperation(operation operations.Operation) LoopCounterBuilder {
	app.operation = operation
	return app
}

// WithIncrement adds an increment to the builder
func (app *loopCounterBuilder) WithIncrement(increment operations.Operation) LoopCounterBuilder {
	app.increment = increment
	return app
}

// Now builds a new LoopCounter instance
func (app *loopCounterBuilder) Now() (LoopCounter, error) {
	if app.assignment == nil {
		return nil, errors.New("the assignment is mandatory in order to build a LoopCounter instance")
	}

	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build a LoopCounter instance")
	}

	if app.increment == nil {
		return nil, errors.New("the increment is mandatory in order to build a LoopCounter instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.assignment.Hash().Bytes(),
		app.operation.Hash().Bytes(),
		app.increment.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLoopCounter(
		*pHash,
		app.assignment,
		app.operation,
		app.increment,
	), nil
}
