package elements

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/constants/tokens/elements/references"
)

type builder struct {
	hashAdapter hash.Adapter
	reference   references.Reference
	rule        string
	constant    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		reference:   nil,
		rule:        "",
		constant:    "",
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

	return createElementWithConstant(*pHash, app.constant), nil
}
