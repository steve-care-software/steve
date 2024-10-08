package elements

import (
	"errors"

	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	token       string
	rule        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		token:       "",
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

// WithToken adds a token to the builder
func (app *builder) WithToken(token string) Builder {
	app.token = token
	return app
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	data := [][]byte{}
	if app.token != "" {
		data = append(data, []byte("token"))
		data = append(data, []byte(app.token))
	}

	if app.rule != "" {
		data = append(data, []byte("rule"))
		data = append(data, []byte(app.rule))
	}

	if len(data) != 2 {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.token != "" {
		return createElementWithToken(*pHash, app.token), nil
	}

	return createElementWithRule(*pHash, app.rule), nil
}
