package suites

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites/expectations"
)

type suiteBuilder struct {
	name        string
	reference   references.Reference
	expectation expectations.Expectation
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:        "",
		reference:   nil,
		expectation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithReference adds a reference to the builder
func (app *suiteBuilder) WithReference(reference references.Reference) SuiteBuilder {
	app.reference = reference
	return app
}

// WithExpectation adds an expectation to the builder
func (app *suiteBuilder) WithExpectation(expectation expectations.Expectation) SuiteBuilder {
	app.expectation = expectation
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.reference == nil {
		return nil, errors.New("the reference is mandatory in order to build a Suite instance")
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mandatory in order to build a Suite instance")
	}

	return createSuite(
		app.name,
		app.reference,
		app.expectation,
	), nil
}
