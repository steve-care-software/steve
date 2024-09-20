package transpiles

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pipelines/transpiles/blocks"
)

type builder struct {
	blocks blocks.Blocks
	root   string
	origin hash.Hash
	target hash.Hash
}

func createBuilder() Builder {
	out := builder{
		blocks: nil,
		root:   "",
		origin: nil,
		target: nil,
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

// WithOrigin adds an origin to the builder
func (app *builder) WithOrigin(origin hash.Hash) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *builder) WithTarget(target hash.Hash) Builder {
	app.target = target
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

	if app.origin == nil {
		return nil, errors.New("the origin hash is mandatory in order to build a Transpile instance")
	}

	if app.target == nil {
		return nil, errors.New("the target hash is mandatory in order to build a Transpile instance")
	}

	return createTranspile(app.blocks, app.root, app.origin, app.target), nil
}
