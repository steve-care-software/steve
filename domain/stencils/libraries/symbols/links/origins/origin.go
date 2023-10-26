package origins

import "github.com/steve-care-software/steve/domain/stencils/pointers"

type origin struct {
	symbol    pointers.Pointer
	direction Direction
}

func createOrigin(
	symbol pointers.Pointer,
) Origin {
	return createOriginInternally(symbol, nil)
}

func createOriginWithDirection(
	symbol pointers.Pointer,
	direction Direction,
) Origin {
	return createOriginInternally(symbol, direction)
}

func createOriginInternally(
	symbol pointers.Pointer,
	direction Direction,
) Origin {
	out := origin{
		symbol:    symbol,
		direction: direction,
	}

	return &out
}

// Symbol returns the symbol
func (obj *origin) Symbol() pointers.Pointer {
	return obj.symbol
}

// HasDirection returns true if there is a direction, false otherwise
func (obj *origin) HasDirection() bool {
	return obj.direction != nil
}

// Direction returns the direction, if any
func (obj *origin) Direction() Direction {
	return obj.direction
}
