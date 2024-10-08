package numerics

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	pFlag       *uint8
	pSize       *uint8
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		pFlag:       nil,
		pSize:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFlag adds a flag to the builder
func (app *builder) WithFlag(flag uint8) Builder {
	app.pFlag = &flag
	return app
}

// WithSize adds a size to the builder
func (app *builder) WithSize(size uint8) Builder {
	app.pSize = &size
	return app
}

// Now builds a new Numeric instance
func (app *builder) Now() (Numeric, error) {
	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a Numeric instance")
	}

	if app.pSize == nil {
		return nil, errors.New("the size is mandatory in order to build a Numeric instance")
	}

	flag := *app.pFlag
	if flag > FlagFloat {
		str := fmt.Sprintf("the flag (%d) is invalid when building the Numeric instance", flag)
		return nil, errors.New(str)
	}

	size := *app.pSize
	if size > Size64 {
		str := fmt.Sprintf("the size (%d) is invalid when building the Numeric instance", size)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		{flag},
		{size},
	})

	if err != nil {
		return nil, err
	}

	return createNumeric(
		*pHash,
		flag,
		size,
	), nil
}
