package expectations

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
)

type expectationBuilder struct {
	references references.References
	isFail     bool
}

func createExpectationBuilder() ExpectationBuilder {
	out := expectationBuilder{
		references: nil,
		isFail:     false,
	}

	return &out
}

// Create initializes the expectationBuilder
func (app *expectationBuilder) Create() ExpectationBuilder {
	return createExpectationBuilder()
}

// WithReferences add references to the expectationBuilder
func (app *expectationBuilder) WithReferences(references references.References) ExpectationBuilder {
	app.references = references
	return app
}

// IsFail flags the expectationBuilder as fail
func (app *expectationBuilder) IsFail() ExpectationBuilder {
	app.isFail = true
	return app
}

// Now builds a new Expectation instance
func (app *expectationBuilder) Now() (Expectation, error) {
	if app.references == nil {
		return nil, errors.New("the references is mandatory in order to build an Expectation instance")
	}

	return createExpectation(
		app.references,
		app.isFail,
	), nil
}
