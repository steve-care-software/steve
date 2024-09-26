package programs

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/initializations"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

type loopCounterBuilder struct {
	hashAdapter    hash.Adapter
	initialization initializations.Initialization
	operation      operations.Operation
	increment      operations.Operation
}

func createLoopCounterBuilder(
	hashAdapter hash.Adapter,
) LoopCounterBuilder {
	out := loopCounterBuilder{
		hashAdapter:    hashAdapter,
		initialization: nil,
		operation:      nil,
		increment:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *loopCounterBuilder) Create() LoopCounterBuilder {
	return createLoopCounterBuilder(
		app.hashAdapter,
	)
}

// WithInitialization adds an initialization to the builder
func (app *loopCounterBuilder) WithInitialization(initialization initializations.Initialization) LoopCounterBuilder {
	app.initialization = initialization
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
	if app.initialization == nil {
		return nil, errors.New("the initialization is mandatory in order to build a LoopCounter instance")
	}

	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build a LoopCounter instance")
	}

	if app.increment == nil {
		return nil, errors.New("the increment is mandatory in order to build a LoopCounter instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.initialization.Hash().Bytes(),
		app.operation.Hash().Bytes(),
		app.increment.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLoopCounter(
		*pHash,
		app.initialization,
		app.operation,
		app.increment,
	), nil
}
