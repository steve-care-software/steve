package suites

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/suites/expectations"
)

type suiteBuilder struct {
	name        string
	link        links.Link
	expectation expectations.Expectations
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:        "",
		link:        nil,
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

// WithLink adds a link to the builder
func (app *suiteBuilder) WithLink(link links.Link) SuiteBuilder {
	app.link = link
	return app
}

// WithExpectations adds an expectation to the builder
func (app *suiteBuilder) WithExpectations(expectation expectations.Expectations) SuiteBuilder {
	app.expectation = expectation
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Suite instance")
	}

	if app.expectation == nil {
		return nil, errors.New("the expectation is mandatory in order to build a Suite instance")
	}

	return createSuite(
		app.name,
		app.link,
		app.expectation,
	), nil
}
