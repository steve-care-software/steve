package origins

import (
	"errors"

	"github.com/steve-care-software/steve/domain/stencils/pointers"
)

type originBuilder struct {
	symbol    pointers.Pointer
	direction Direction
}

func createOriginBuilder() OriginBuilder {
	out := originBuilder{
		symbol:    nil,
		direction: nil,
	}

	return &out
}

// Create initializes the builder
func (app *originBuilder) Create() OriginBuilder {
	return createOriginBuilder()
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

	if app.direction != nil {
		return createOriginWithDirection(app.symbol, app.direction), nil
	}

	return createOrigin(app.symbol), nil
}
