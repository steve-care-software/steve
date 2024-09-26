package programs

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	input       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		input:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to builld a Program instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to builld a Program instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte(app.input),
	})

	if err != nil {
		return nil, err
	}

	return createProgram(
		*pHash,
		app.name,
		app.input,
	), nil
}
