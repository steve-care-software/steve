package blocks

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks/lines"
	"github.com/steve-care-software/steve/hash"
)

type blockBuilder struct {
	hashAdapter hash.Adapter
	name        string
	lines       lines.Lines
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		lines:       nil,
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

// Now builds a new Block instance
func (app *blockBuilder) Now() (Block, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Block instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines are mandatory in order to build a Block instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		app.lines.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createBlock(*pHash, app.name, app.lines), nil
}
