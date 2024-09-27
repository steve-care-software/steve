package targets

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
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

// Now builds a new Target instance
func (app *builder) Now() (Target, error) {
	data := [][]byte{}
	if app.constant != "" {
		data = append(data, []byte("constant"))
		data = append(data, []byte(app.constant))
	}

	if app.rule != "" {
		data = append(data, []byte("rule"))
		data = append(data, []byte(app.rule))
	}

	if len(data) != 2 {
		return nil, errors.New("the Target is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.constant != "" {
		return createTargetWithConstant(*pHash, app.constant), nil
	}

	return createTargetWithRule(*pHash, app.rule), nil
}
