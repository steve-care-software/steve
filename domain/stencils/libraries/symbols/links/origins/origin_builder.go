package origins

import (
	"errors"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

type originBuilder struct {
	hashAdapter hash.Adapter
	symbol      pointers.Pointer
	direction   Direction
}

func createOriginBuilder(
	hashAdapter hash.Adapter,
) OriginBuilder {
	out := originBuilder{
		hashAdapter: hashAdapter,
		symbol:      nil,
		direction:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *originBuilder) Create() OriginBuilder {
	return createOriginBuilder(
		app.hashAdapter,
	)
}

// WithSymbol adds symbol to the builder
func (app *originBuilder) WithSymbol(symbol pointers.Pointer) OriginBuilder {
	app.symbol = symbol
	return app
}

// WithDirection adds direction to the builder
func (app *originBuilder) WithDirection(direction Direction) OriginBuilder {
	app.direction = direction
	return app
}

// Now builds a new Origin instance
func (app *originBuilder) Now() (Origin, error) {
	if app.symbol == nil {
		return nil, errors.New("the symbol is mandatory in order to build an Origin instance")
	}

	data := [][]byte{
		app.symbol.Hash().Bytes(),
	}

	if app.direction != nil {
		data = append(data, app.direction.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.direction != nil {
		return createOriginWithDirection(*pHash, app.symbol, app.direction), nil
	}

	return createOrigin(*pHash, app.symbol), nil
}
