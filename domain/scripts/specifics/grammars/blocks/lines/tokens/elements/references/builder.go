package references

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	grammar     string
	block       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		grammar:     "",
		block:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar string) Builder {
	app.grammar = grammar
	return app
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block string) Builder {
	app.block = block
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.grammar == "" {
		return nil, errors.New("the grammar is mandatory in order to build a Reference instance")
	}

	if app.block == "" {
		return nil, errors.New("the block is mandatory in order to build a Reference instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.grammar),
		[]byte(app.block),
	})

	if err != nil {
		return nil, err
	}

	return createReference(
		*pHash,
		app.grammar,
		app.block,
	), nil
}
