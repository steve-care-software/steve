package operations

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/steve/hash"
)

type singleSwordBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	pFlag       *uint8
}

func createSingleSwordBuilder(
	hashAdapter hash.Adapter,
) SingleSwordBuilder {
	out := singleSwordBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		pFlag:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *singleSwordBuilder) Create() SingleSwordBuilder {
	return createSingleSwordBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *singleSwordBuilder) WithVariable(variable string) SingleSwordBuilder {
	app.variable = variable
	return app
}

// WithFlag adds a flag to the builder
func (app *singleSwordBuilder) WithFlag(flag uint8) SingleSwordBuilder {
	app.pFlag = &flag
	return app
}

// Now builds a new SingleSword instance
func (app *singleSwordBuilder) Now() (SingleSword, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a SingleSword instance")
	}

	if app.pFlag == nil {
		return nil, errors.New("the flag is mandatory in order to build a SingleSword instance")
	}

	flag := *app.pFlag
	if flag > SingleSwordMinus {
		str := fmt.Sprintf("the flag (%d) is invalid when building the SingleSword instance", flag)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		{flag},
	})

	if err != nil {
		return nil, err
	}

	return createSingleSword(
		*pHash,
		app.variable,
		flag,
	), nil
}
