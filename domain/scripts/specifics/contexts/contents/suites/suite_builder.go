package suites

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/domain/hash"
)

type suiteBuilder struct {
	hashAdapter hash.Adapter
	name        string
	path        []string
	isFail      bool
}

func createSuiteBuilder(
	hashAdapter hash.Adapter,
) SuiteBuilder {
	out := suiteBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		path:        nil,
		isFail:      false,
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

// WithPath adds a path to the builder
func (app *suiteBuilder) WithPath(path []string) SuiteBuilder {
	app.path = path
	return app
}

// IsFail flags the builder as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Suite instance")
	}

	isFail := "false"
	if app.isFail {
		isFail = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte(strings.Join(app.path, ",")),
		[]byte(isFail),
	})

	if err != nil {
		return nil, err
	}

	return createSuite(*pHash, app.name, app.path, app.isFail), nil
}
