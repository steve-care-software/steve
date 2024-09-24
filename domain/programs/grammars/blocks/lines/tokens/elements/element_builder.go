package elements

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	rule        string
	block       string
	spacer      string
	constant    string
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		rule:        "",
		block:       "",
		spacer:      "",
		constant:    "",
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithRule adds a rule to the builder
func (app *elementBuilder) WithRule(rule string) ElementBuilder {
	app.rule = rule
	return app
}

// WithBlock adds a block to the builder
func (app *elementBuilder) WithBlock(block string) ElementBuilder {
	app.block = block
	return app
}

// WithSpacer adds a spacer to the builder
func (app *elementBuilder) WithSpacer(spacer string) ElementBuilder {
	app.spacer = spacer
	return app
}

// WithConstant adds a constant to the builder
func (app *elementBuilder) WithConstant(constant string) ElementBuilder {
	app.constant = constant
	return app
}

// Now builds a new Element
func (app *elementBuilder) Now() (Element, error) {
	data := [][]byte{}
	if app.rule != "" {
		data = append(data, []byte(app.rule))
	}

	if app.block != "" {
		data = append(data, []byte(app.block))
	}

	if app.spacer != "" {
		data = append(data, []byte(app.spacer))
	}

	if app.constant != "" {
		data = append(data, []byte(app.constant))
	}

	if len(data) != 1 {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.rule != "" {
		return createElementWithRule(*pHash, app.rule), nil
	}

	if app.block != "" {
		return createElementWithBlock(*pHash, app.block), nil
	}

	if app.spacer != "" {
		return createElementWithSpacer(*pHash, app.spacer), nil
	}

	return createElementWithConstant(*pHash, app.constant), nil
}
