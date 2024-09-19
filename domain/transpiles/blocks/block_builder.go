package blocks

import (
	"errors"

	"github.com/steve-care-software/steve/domain/transpiles/blocks/lines"
)

type blockBuilder struct {
	name  string
	lines lines.Lines
}

func createBlockBuilder() BlockBuilder {
	out := blockBuilder{
		name:  "",
		lines: nil,
	}

	return &out
}

// Create initializes the builder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder()
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

	return createBlock(app.name, app.lines), nil
}
