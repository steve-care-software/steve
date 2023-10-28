package links

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/origins"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/preparations"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/suites"
)

type builder struct {
	hashAdapter  hash.Adapter
	origins      origins.Origins
	execution    executions.Execution
	preparations preparations.Preparations
	suites       suites.Suites
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		origins:      nil,
		execution:    nil,
		preparations: nil,
		suites:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOrigins add origins to the builder
func (app *builder) WithOrigins(origins origins.Origins) Builder {
	app.origins = origins
	return app
}

// WithExecution add execution to the builder
func (app *builder) WithExecution(execution executions.Execution) Builder {
	app.execution = execution
	return app
}

// WithPreparations add preparations to the builder
func (app *builder) WithPreparations(preparations preparations.Preparations) Builder {
	app.preparations = preparations
	return app
}

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites suites.Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.origins == nil {
		return nil, errors.New("the origins is mandatory in order to build a Link instance")
	}

	if app.execution == nil {
		return nil, errors.New("the execution is mandatory in order to build a Link instance")
	}

	if app.preparations == nil {
		return nil, errors.New("the preparations is mandatory in order to build a Link instance")
	}

	data := [][]byte{
		app.origins.Hash().Bytes(),
		app.execution.Hash().Bytes(),
		app.preparations.Hash().Bytes(),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createLinkWithSuites(*pHash, app.origins, app.execution, app.preparations, app.suites), nil
	}

	return createLink(*pHash, app.origins, app.execution, app.preparations), nil
}
