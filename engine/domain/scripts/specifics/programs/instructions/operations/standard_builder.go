package operations

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/commons/hash"
)

type standardBuilder struct {
	hashAdapter hash.Adapter
	first       Operation
	second      Operation
	pFlag       *uint8
}

func createStandardBuilder(
	hashAdapter hash.Adapter,
) StandardBuilder {
	out := standardBuilder{
		hashAdapter: hashAdapter,
		first:       nil,
		second:      nil,
		pFlag:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *standardBuilder) Create() StandardBuilder {
	return createStandardBuilder(
		app.hashAdapter,
	)
}

// WithFirst adds a first operation to the builder
func (app *standardBuilder) WithFirst(first Operation) StandardBuilder {
	app.first = first
	return app
}

// WithSecond adds a second operation to the builder
func (app *standardBuilder) WithSecond(second Operation) StandardBuilder {
	app.second = second
	return app
}

// WithFlag adds a flag to the builder
func (app *standardBuilder) WithFlag(flag uint8) StandardBuilder {
	app.pFlag = &flag
	return app
}

// Now builds a new Standard instance
func (app *standardBuilder) Now() (Standard, error) {
	if app.first == nil {
		return nil, errors.New("the first Operation is mandatory in order to build a Standard instance")
	}

	if app.second == nil {
		return nil, errors.New("the second Operation is mandatory in order to build a Standard instance")
	}

	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Standard instance")
	}

	flag := *app.pFlag
	if flag > StandardXor {
		str := fmt.Sprintf("the flag (%d) is invalid when building the Standard instance", flag)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.first.Hash().Bytes(),
		app.second.Hash().Bytes(),
		{flag},
	})

	if err != nil {
		return nil, err
	}

	return createStandard(
		*pHash,
		app.first,
		app.second,
		flag,
	), nil
}
