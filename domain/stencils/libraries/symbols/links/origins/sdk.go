package origins

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewOriginBuilder creates a new origin builder
func NewOriginBuilder() OriginBuilder {
	hashAdapter := hash.NewAdapter()
	return createOriginBuilder(
		hashAdapter,
	)
}

// NewDirectionBuilder creates a new direction builder
func NewDirectionBuilder() DirectionBuilder {
	return createDirectionBuilder()
}

// Builder represents an origin builder
type Builder interface {
	Create() Builder
	WithList(list []Origin) Builder
	Now() (Origins, error)
}

// Origins represents origins
type Origins interface {
	Hash() hash.Hash
	List() []Origin
}

// OriginBuilder represents an origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithSymbol(symbol pointers.Pointer) OriginBuilder
	WithDirection(direction Direction) OriginBuilder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	Symbol() pointers.Pointer
	HasDirection() bool
	Direction() Direction
}

// DirectionBuilder represents a direction builder
type DirectionBuilder interface {
	Create() DirectionBuilder
	WithNext(next Origin) DirectionBuilder
	WithPrevious(previous Origin) DirectionBuilder
	Now() (Direction, error)
}

// Direction represents a direction
type Direction interface {
	Hash() hash.Hash
	IsNext() bool
	Next() Origin
	IsPrevious() bool
	Previous() Origin
}
