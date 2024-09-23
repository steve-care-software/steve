package elements

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens/elements/references"
)

type builder struct {
	hashAdapter hash.Adapter
	reference   references.Reference
	rule        string
	constant    string
	block       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		reference:   nil,
		rule:        "",
		constant:    "",
		block:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithReference adds a reference to the builder
func (app *builder) WithReference(reference references.Reference) Builder {
	app.reference = reference
	return app
}

// WithRule adds a rule to the builder
func (app *builder) WithRule(rule string) Builder {
	app.rule = rule
	return app
}

// WithConstant adds a constant to the builder
func (app *builder) WithConstant(constant string) Builder {
	app.constant = constant
	return app
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block string) Builder {
	app.block = block
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	data := [][]byte{}
	if app.reference != nil {
		data = append(data, app.reference.Hash().Bytes())
	}

	if app.rule != "" {
		data = append(data, []byte(app.rule))
	}

	if app.constant != "" {
		data = append(data, []byte(app.constant))
	}

	if app.block != "" {
		data = append(data, []byte(app.block))
	}

	if len(data) != 1 {
		return nil, errors.New("the Element is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.reference != nil {
		return createElementWithReference(*pHash, app.reference), nil
	}

	if app.rule != "" {
		return createElementWithRule(*pHash, app.rule), nil
	}

	if app.constant != "" {
		return createElementWithConstant(*pHash, app.constant), nil
	}

	return createElementWithBlock(*pHash, app.block), nil
}
