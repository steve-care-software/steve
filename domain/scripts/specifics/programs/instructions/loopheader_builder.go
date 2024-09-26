package instructions

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
)

type loopHeaderBuilder struct {
	hashAdapter hash.Adapter
	counter     LoopCounter
	keyValue    LoopKeyValue
	isInfinite  bool
}

func createLoopHeaderBuilder(
	hashAdapter hash.Adapter,
) LoopHeaderBuilder {
	out := loopHeaderBuilder{
		hashAdapter: hashAdapter,
		counter:     nil,
		keyValue:    nil,
		isInfinite:  false,
	}

	return &out
}

// Create initializes the builder
func (app *loopHeaderBuilder) Create() LoopHeaderBuilder {
	return createLoopHeaderBuilder(
		app.hashAdapter,
	)
}

// WithCounter adds a counter to the builder
func (app *loopHeaderBuilder) WithCounter(counter LoopCounter) LoopHeaderBuilder {
	app.counter = counter
	return app
}

// WithKeyValue adds a keyValue to the builder
func (app *loopHeaderBuilder) WithKeyValue(keyValue LoopKeyValue) LoopHeaderBuilder {
	app.keyValue = keyValue
	return app
}

// IsInfinite flags the builder as infinite
func (app *loopHeaderBuilder) IsInfinite() LoopHeaderBuilder {
	app.isInfinite = true
	return app
}

// Now builds a new LoopHeader instance
func (app *loopHeaderBuilder) Now() (LoopHeader, error) {
	data := [][]byte{}
	if app.counter != nil {
		data = append(data, app.counter.Hash().Bytes())
	}

	if app.keyValue != nil {
		data = append(data, app.keyValue.Hash().Bytes())
	}

	if app.isInfinite {
		data = append(data, []byte("infinite"))
	}

	if len(data) != 1 {
		return nil, errors.New("the LoopHeader is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.counter != nil {
		return createLoopHeaderWithCounter(*pHash, app.counter), nil
	}

	if app.keyValue != nil {
		return createLoopHeaderWithKeyValue(*pHash, app.keyValue), nil
	}

	return createLoopHeaderWithKeyInfinite(*pHash), nil
}
