package values

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	constant    string
	rule        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
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

	if app.constant != "" {
		return createValueWithConstant(*pHash, app.constant), nil
	}

	return createValueWithRule(*pHash, app.rule), nil
}
