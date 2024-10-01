package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type loopInstructionsBuilder struct {
	hashAdapter hash.Adapter
	list        []LoopInstruction
}

func createLoopInstructionsBuilder(
	hashAdapter hash.Adapter,
) LoopInstructionsBuilder {
	out := loopInstructionsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the loopInstructionsBuilder
func (app *loopInstructionsBuilder) Create() LoopInstructionsBuilder {
	return createLoopInstructionsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the loopInstructionsBuilder
func (app *loopInstructionsBuilder) WithList(list []LoopInstruction) LoopInstructionsBuilder {
	app.list = list
	return app
}

// Now builds a new LoopInstructions instance
func (app *loopInstructionsBuilder) Now() (LoopInstructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 LoopInstruction in order to build a LoopInstructions instance")
	}

	data := [][]byte{}
	for _, oneLoopInstruction := range app.list {
		data = append(data, oneLoopInstruction.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createLoopInstructions(
		*pHash,
		app.list,
	), nil
}
