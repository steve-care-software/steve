package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/instructions/operations"
)

type conditionBuilder struct {
	hashAdapter  hash.Adapter
	operation    operations.Operation
	instructions Instructions
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter:  hashAdapter,
		operation:    nil,
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithOperation adds an operation to the builder
func (app *conditionBuilder) WithOperation(operation operations.Operation) ConditionBuilder {
	app.operation = operation
	return app
}

// WithInstructions adds an instructions to the builder
func (app *conditionBuilder) WithInstructions(instructions Instructions) ConditionBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build a Condition instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.operation.Hash().Bytes(),
		app.instructions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(
		*pHash,
		app.operation,
		app.instructions,
	), nil
}
