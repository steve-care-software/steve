package origins

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

type origin struct {
	hash      hash.Hash
	symbol    pointers.Pointer
	direction Direction
}

func createOrigin(
	hash hash.Hash,
	symbol pointers.Pointer,
) Origin {
	return createOriginInternally(hash, symbol, nil)
}

func createOriginWithDirection(
	hash hash.Hash,
	symbol pointers.Pointer,
	direction Direction,
) Origin {
	return createOriginInternally(hash, symbol, direction)
}

func createOriginInternally(
	hash hash.Hash,
	symbol pointers.Pointer,
	direction Direction,
) Origin {
	out := origin{
		hash:      hash,
		symbol:    symbol,
		direction: direction,
	}

	return &out
}

// Hash returns the hash
func (obj *origin) Hash() hash.Hash {
	return obj.hash
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
