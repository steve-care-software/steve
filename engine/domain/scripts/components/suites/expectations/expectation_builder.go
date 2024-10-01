package expectations

import (
	"errors"
	"strings"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type expectationBuilder struct {
	hashAdapter hash.Adapter
	path        []string
	isFail      bool
}

func createExpectationBuilder(
	hashAdapter hash.Adapter,
) ExpectationBuilder {
	out := expectationBuilder{
		hashAdapter: hashAdapter,
		path:        nil,
		isFail:      false,
	}

	return &out
}

// Create initializes the builder
func (app *expectationBuilder) Create() ExpectationBuilder {
	return createExpectationBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *expectationBuilder) WithPath(path []string) ExpectationBuilder {
	app.path = path
	return app
}

// IsFail flags the builder as fail
func (app *expectationBuilder) IsFail() ExpectationBuilder {
	app.isFail = true
	return app
}

// Now builds a new Expectation instance
func (app *expectationBuilder) Now() (Expectation, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("there must be at least 1 path element in order to build an Expectation instance")
	}

	isFail := "false"
	if app.isFail {
		isFail = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strings.Join(app.path, ",")),
		[]byte(isFail),
	})

	if err != nil {
		return nil, err
	}

	return createExpectation(
		*pHash,
		app.path,
		app.isFail,
	), nil
}
