package blocks

import (
	"errors"

	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents"
)

type blockBuilder struct {
	hashAdapter hash.Adapter
	content     contents.Content
	result      []byte
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		result:      nil,
	}

	return &out
}

// Create initializes the blockBuilder
func (app *blockBuilder) Create() BlockBuilder {
	return createBlockBuilder(
		app.hashAdapter,
	)
}

// WithContent adds content to the blockBuilder
func (app *blockBuilder) WithContent(content contents.Content) BlockBuilder {
	app.content = content
	return app
}

// WithResult adds result to the blockBuilder
func (app *blockBuilder) WithResult(result []byte) BlockBuilder {
	app.result = result
	return app
}

// Now builds a new BLock instance
func (app *blockBuilder) Now() (Block, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Block instance")
	}

	if app.result != nil && len(app.result) <= 0 {
		app.result = nil
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a BLock instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		app.result,
	})

	if err != nil {
		return nil, err
	}

	return createBlock(
		*pHash,
		app.content,
		app.result,
	), nil
}
