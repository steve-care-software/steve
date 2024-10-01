package values

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	block       string
	constant    string
	rule        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		block:       "",
		constant:    "",
		rule:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block string) Builder {
	app.block = block
	return app
}

// WithConstant adds a constant to the builder
func (app *builder) WithConstant(constant string) Builder {
	app.constant = constant
	return app
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	data := [][]byte{}
	if app.block != "" {
		data = append(data, []byte(app.block))
	}

	if app.constant != "" {
		data = append(data, []byte(app.constant))
	}

	if app.rule != "" {
		data = append(data, []byte(app.rule))
	}

	if len(data) != 1 {
		return nil, errors.New("the Value is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.block != "" {
		return createValueWithBlock(*pHash, app.block), nil
	}

	if app.constant != "" {
		return createValueWithConstant(*pHash, app.constant), nil
	}

	return createValueWithRule(*pHash, app.rule), nil
}
