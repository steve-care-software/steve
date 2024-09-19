package transpiles

import (
	"errors"

	"github.com/steve-care-software/steve/domain/transpiles/blocks"
)

type builder struct {
	blocks blocks.Blocks
	root   string
}

func createBuilder() Builder {
	out := builder{
		blocks: nil,
		root:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBlocks add blocks to the builder
func (app *builder) WithBlocks(blocks blocks.Blocks) Builder {
	app.blocks = blocks
	return app
}

// WithRoot add root to the builder
func (app *builder) WithRoot(root string) Builder {
	app.root = root
	return app
}

// Now builds a new Transpile instance
func (app *builder) Now() (Transpile, error) {
	if app.blocks == nil {
		return nil, errors.New("the blocks is mandatory in order to build a Transpile instance")
	}

	if app.root == "" {
		return nil, errors.New("the root is mandatory in order to build a Transpile instance")
	}

	return createTranspile(app.blocks, app.root), nil
}
