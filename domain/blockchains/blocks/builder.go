package blocks

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	content     contents.Content
	result      []byte
	difficulty  uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		result:      nil,
		difficulty:  0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content contents.Content) Builder {
	app.content = content
	return app
}

// WithResult adds result to the builder
func (app *builder) WithResult(result []byte) Builder {
	app.result = result
	return app
}

// WithDifficulty adds a difficulty to the builder
func (app *builder) WithDifficulty(difficulty uint) Builder {
	app.difficulty = difficulty
	return app
}

// Now builds a new BLock instance
func (app *builder) Now() (Block, error) {
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
