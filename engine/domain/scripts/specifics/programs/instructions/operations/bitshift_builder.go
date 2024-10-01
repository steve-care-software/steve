package operations

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type bitshiftBuilder struct {
	hashAdapter hash.Adapter
	operation   Operation
	pFlag       *uint8
	amount      uint8
}

func createBitshiftBuilder(
	hashAdapter hash.Adapter,
) BitShiftBuilder {
	out := bitshiftBuilder{
		hashAdapter: hashAdapter,
		operation:   nil,
		pFlag:       nil,
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *bitshiftBuilder) Create() BitShiftBuilder {
	return createBitshiftBuilder(
		app.hashAdapter,
	)
}

// WithOperation adds an operation to the builder
func (app *bitshiftBuilder) WithOperation(operation Operation) BitShiftBuilder {
	app.operation = operation
	return app
}

// WithFlag adds a flag to the builder
func (app *bitshiftBuilder) WithFlag(flag uint8) BitShiftBuilder {
	app.pFlag = &flag
	return app
}

// WithAmount adds an amount to the builder
func (app *bitshiftBuilder) WithAmount(amount uint8) BitShiftBuilder {
	app.amount = amount
	return app
}

// Now builds a new BitShift instance
func (app *bitshiftBuilder) Now() (BitShift, error) {
	if app.operation == nil {
		return nil, errors.New("the operation is mandatory in order to build a BitShift instance")
	}

	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a BitShift instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a BitShift instance")
	}

	flag := *app.pFlag
	if flag > BitShiftRight {
		str := fmt.Sprintf("the flag (%d) is invalid when building the BitShift instance", flag)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.operation.Hash().Bytes(),
		{flag},
		{app.amount},
	})

	if err != nil {
		return nil, err
	}

	return createBitshift(
		*pHash,
		app.operation,
		flag,
		app.amount,
	), nil
}
