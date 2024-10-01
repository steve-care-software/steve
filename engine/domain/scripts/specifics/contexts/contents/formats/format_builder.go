package formats

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents/formats/suites"
)

type formatBuilder struct {
	hashAdapter hash.Adapter
	point       string
	grammar     string
	suites      suites.Suites
}

func createFormatBuilder(
	hashAdapter hash.Adapter,
) FormatBuilder {
	out := formatBuilder{
		hashAdapter: hashAdapter,
		point:       "",
		grammar:     "",
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *formatBuilder) Create() FormatBuilder {
	return createFormatBuilder(
		app.hashAdapter,
	)
}

// WithPoint adds a point to the builder
func (app *formatBuilder) WithPoint(point string) FormatBuilder {
	app.point = point
	return app
}

// WithGrammar adds a grammar to the builder
func (app *formatBuilder) WithGrammar(grammar string) FormatBuilder {
	app.grammar = grammar
	return app
}

// WithSuites adds a suites to the builder
func (app *formatBuilder) WithSuites(suites suites.Suites) FormatBuilder {
	app.suites = suites
	return app
}

// Now builds a Format instance
func (app *formatBuilder) Now() (Format, error) {
	if app.point == "" {
		return nil, errors.New("the point is mandatory in order to build a Format instance")
	}

	if app.grammar == "" {
		return nil, errors.New("the grammar is mandatory in order to build a Format instance")
	}

	data := [][]byte{
		[]byte(app.point),
		[]byte(app.grammar),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createFormatWithSuites(
			*pHash,
			app.point,
			app.grammar,
			app.suites,
		), nil
	}

	return createFormat(
		*pHash,
		app.point,
		app.grammar,
	), nil
}
