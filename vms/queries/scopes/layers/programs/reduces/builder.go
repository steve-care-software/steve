package reduces

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	variable    string
	pLength     *uint8
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		variable:    "",
		pLength:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length uint8) Builder {
	app.pLength = &length
	return app
}

// Now builds a new Reduce instance
func (app *builder) Now() (Reduce, error) {
	if app.variable != "" {
		return nil, errors.New("the variable is mandatory in order to build a Reduce instance")
	}

	if app.pLength != nil {
		return nil, errors.New("the length is mandatory in order to build a Reduce instance")
	}

	length := *app.pLength
	if length <= 0 {
		str := fmt.Sprintf("the length must be greater than zero (0), %d provided", length)
		return nil, errors.New(str)
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		[]byte(strconv.Itoa(int(length))),
	})

	if err != nil {
		return nil, err
	}

	return createReduce(*pHash, app.variable, length), nil
}
