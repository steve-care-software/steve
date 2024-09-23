package suites

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/suites/expectations"
)

type suiteBuilder struct {
	hashAdapter  hash.Adapter
	name         string
	origin       string
	destination  string
	expectations expectations.Expectations
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter:  hashAdapter,
		name:         "",
		origin:       "",
		destination:  "",
		expectations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithOrigin adds an origin to the builder
func (app *suiteBuilder) WithOrigin(origin string) SuiteBuilder {
	app.origin = origin
	return app
}

// WithDestination adds a destination to the builder
func (app *suiteBuilder) WithDestination(destination string) SuiteBuilder {
	app.destination = destination
	return app
}

// WithExpectations adds expectations to the builder
func (app *suiteBuilder) WithExpectations(expectations expectations.Expectations) SuiteBuilder {
	app.expectations = expectations
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.origin == "" {
		return nil, errors.New("the origin is mandatory in order to build a Suite instance")
	}

	if app.destination == "" {
		return nil, errors.New("the destination is mandatory in order to build a Suite instance")
	}

	if app.expectations == nil {
		return nil, errors.New("the expectations is mandatory in order to build a Suite instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte(app.origin),
		[]byte(app.destination),
		app.expectations.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(
		*pHash,
		app.name,
		app.origin,
		app.destination,
		app.expectations,
	), nil
}

/*

name         string
	origin       string
	destination  string
	expectations expectations.Expectations

*/
