package expectations

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
)

type builder struct {
	references references.References
	isFail     bool
}

func createBuilder() Builder {
	out := builder{
		references: nil,
		isFail:     false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReferences add references to the builder
func (app *builder) WithReferences(references references.References) Builder {
	app.references = references
	return app
}

// IsFail flags the builder as fail
func (app *builder) IsFail() Builder {
	app.isFail = true
	return app
}

// Now builds a new Expectation instance
func (app *builder) Now() (Expectation, error) {
	if app.references == nil {
		return nil, errors.New("the references is mandatory in order to build an Expectation instance")
	}

	return createExpectation(
		app.references,
		app.isFail,
	), nil
}
