package blocks

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

type blockBuilder struct {
	hashAdapter hash.Adapter
	content     contents.Content
	result      []byte
	difficulty  uint8
}

func createBlockBuilder(
	hashAdapter hash.Adapter,
) BlockBuilder {
	out := blockBuilder{
		hashAdapter: hashAdapter,
		content:     nil,
		result:      nil,
		difficulty:  0,
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

// WithDifficulty adds a difficulty to the blockBuilder
func (app *blockBuilder) WithDifficulty(difficulty uint8) BlockBuilder {
	app.difficulty = difficulty
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

	if app.difficulty <= 0 {
		return nil, errors.New("the difficulty must be greater than zero (0) in order to build a Block instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		app.result,
		[]byte(strconv.Itoa(int(app.difficulty))),
	})

	if err != nil {
		return nil, err
	}

	return createBlock(
		*pHash,
		app.content,
		app.result,
		app.difficulty,
	), nil
}
