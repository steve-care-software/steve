package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type loopBuilder struct {
	hashAdapter  hash.Adapter
	header       LoopHeader
	instructions LoopInstructions
}

func createLoopBuilder(
	hashAdapter hash.Adapter,
) LoopBuilder {
	out := loopBuilder{
		hashAdapter:  hashAdapter,
		header:       nil,
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *loopBuilder) Create() LoopBuilder {
	return createLoopBuilder(
		app.hashAdapter,
	)
}

// WithHeader adds an header to the builder
func (app *loopBuilder) WithHeader(header LoopHeader) LoopBuilder {
	app.header = header
	return app
}

// WithInstructions adds an instructions to the builder
func (app *loopBuilder) WithInstructions(instructions LoopInstructions) LoopBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Loop instance
func (app *loopBuilder) Now() (Loop, error) {
	if app.header == nil {
		return nil, errors.New("the header is mandatory in order to build a Loop instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Loop instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.header.Hash().Bytes(),
		app.instructions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLoop(
		*pHash,
		app.header,
		app.instructions,
	), nil
}
