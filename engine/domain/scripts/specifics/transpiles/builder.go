package transpiles

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles/blocks"
)

type builder struct {
	hashAdapter hash.Adapter
	blocks      blocks.Blocks
	root        string
	origin      string
	target      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		blocks:      nil,
		root:        "",
		origin:      "",
		target:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
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
func (app *builder) WithOrigin(origin string) Builder {
	app.origin = origin
	return app
}

// WithTarget adds a target to the builder
func (app *builder) WithTarget(target string) Builder {
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

	if app.origin == "" {
		return nil, errors.New("the origin hash is mandatory in order to build a Transpile instance")
	}

	if app.target == "" {
		return nil, errors.New("the target hash is mandatory in order to build a Transpile instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.blocks.Hash().Bytes(),
		[]byte(app.root),
		[]byte(app.origin),
		[]byte(app.target),
	})

	if err != nil {
		return nil, err
	}

	return createTranspile(*pHash, app.blocks, app.root, app.origin, app.target), nil
}
