package blocks

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines"
	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/suites"
)

type blockBuilder struct {
	hashAdapter hash.Adapter
	name        string
	lines       lines.Lines
	suites      suites.Suites
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		lines:       nil,
		suites:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *blockBuilder) WithName(name string) BlockBuilder {
	app.name = name
	return app
}

// WithLines add lines to the builder
func (app *blockBuilder) WithLines(lines lines.Lines) BlockBuilder {
	app.lines = lines
	return app
}

// WithSuites add suites to the builder
func (app *blockBuilder) WithSuites(suites suites.Suites) BlockBuilder {
	app.suites = suites
	return app
}

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Block instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Block instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.lines.Hash().Bytes(),
	}

	if app.suites != nil {
		data = append(data, app.suites.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.suites != nil {
		return createBlockWithSuites(*pHash, app.name, app.lines, app.suites), nil
	}

	return createBlock(*pHash, app.name, app.lines), nil
}
