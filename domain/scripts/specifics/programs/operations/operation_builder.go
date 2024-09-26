package operations

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/values"
)

type operationBuilder struct {
	hashAdapter hash.Adapter
	standard    Standard
	singleSword SingleSword
	bitshift    BitShift
	value       values.Value
}

func createOperationBuilder(
	hashAdapter hash.Adapter,
) OperationBuilder {
	out := operationBuilder{
		hashAdapter: hashAdapter,
		standard:    nil,
		singleSword: nil,
		bitshift:    nil,
		value:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *operationBuilder) Create() OperationBuilder {
	return createOperationBuilder(
		app.hashAdapter,
	)
}

// WithStandard adds standard to the builder
func (app *operationBuilder) WithStandard(standard Standard) OperationBuilder {
	app.standard = standard
	return app
}

// WithSingleSword adds singleSword to the builder
func (app *operationBuilder) WithSingleSword(singleSword SingleSword) OperationBuilder {
	app.singleSword = singleSword
	return app
}

// WithBitShift adds bitshift to the builder
func (app *operationBuilder) WithBitShift(bitshift BitShift) OperationBuilder {
	app.bitshift = bitshift
	return app
}

// WithValue adds value to the builder
func (app *operationBuilder) WithValue(value values.Value) OperationBuilder {
	app.value = value
	return app
}

// Now builds a new Operation instance
func (app *operationBuilder) Now() (Operation, error) {
	data := [][]byte{}
	if app.standard != nil {
		data = append(data, app.standard.Hash().Bytes())
	}

	if app.singleSword != nil {
		data = append(data, app.singleSword.Hash().Bytes())
	}

	if app.bitshift != nil {
		data = append(data, app.bitshift.Hash().Bytes())
	}

	if app.value != nil {
		data = append(data, app.value.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Operation is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.standard != nil {
		return createOperationWithStandard(
			*pHash,
			app.standard,
		), nil
	}

	if app.singleSword != nil {
		return createOperationWithSingleSword(
			*pHash,
			app.singleSword,
		), nil
	}

	if app.bitshift != nil {
		return createOperationWithBitShift(
			*pHash,
			app.bitshift,
		), nil
	}

	return createOperationWithValue(
		*pHash,
		app.value,
	), nil
}
