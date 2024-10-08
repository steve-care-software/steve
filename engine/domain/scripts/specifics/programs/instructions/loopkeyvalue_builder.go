package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/operations"
	"github.com/steve-care-software/steve/hash"
)

type loopKeyValueBuilder struct {
	hashAdapter hash.Adapter
	keyName     string
	valueName   string
	operation   operations.Operation
}

func createLoopKeyValueBuilder(
	hashAdapter hash.Adapter,
) LoopKeyValueBuilder {
	out := loopKeyValueBuilder{
		hashAdapter: hashAdapter,
		keyName:     "",
		valueName:   "",
		operation:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *loopKeyValueBuilder) Create() LoopKeyValueBuilder {
	return createLoopKeyValueBuilder(
		app.hashAdapter,
	)
}

// WithKeyname adds a keyname to the builder
func (app *loopKeyValueBuilder) WithKeyname(keyname string) LoopKeyValueBuilder {
	app.keyName = keyname
	return app
}

// WithValueName adds a valueName to the builder
func (app *loopKeyValueBuilder) WithValueName(valueName string) LoopKeyValueBuilder {
	app.valueName = valueName
	return app
}

// WithOperation adds an operation to the builder
func (app *loopKeyValueBuilder) WithOperation(operation operations.Operation) LoopKeyValueBuilder {
	app.operation = operation
	return app
}

// Now builds a new LoopKeyValue instance
func (app *loopKeyValueBuilder) Now() (LoopKeyValue, error) {
	if app.keyName == "" {
		return nil, errors.New("the keyName is mandatory in order to build a LoopKeyValue instance")
	}

	if app.valueName == "" {
		return nil, errors.New("the valueName is mandatory in order to build a LoopKeyValue instance")
	}

	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build a LoopKeyValue instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.keyName),
		[]byte(app.valueName),
		app.operation.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLoopKeyValue(
		*pHash,
		app.keyName,
		app.valueName,
		app.operation,
	), nil
}
