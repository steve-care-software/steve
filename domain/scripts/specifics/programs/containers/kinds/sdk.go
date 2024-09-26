package kinds

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers/kinds/numerics"
)

const (
	// EnginePath represents the engine path
	EnginePath (uint8) = iota

	// EngineRoute represents the engine route
	EngineRoute
)

const (
	// RemainingString represents the remaining string
	RemainingString (uint8) = iota

	// RemainingBool represents the remaining bool
	RemainingBool
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the kind builder
type Builder interface {
	Create() Builder
	WithNumeric(numeric numerics.Numeric) Builder
	WithEngine(engine uint8) Builder
	WithRemaining(remaining uint8) Builder
	Now() (Kind, error)
}

// Kind represents a kind
type Kind interface {
	Hash() hash.Hash
	IsNumeric() bool
	Numeric() numerics.Numeric
	IsEngine() bool
	Engine() *uint8 // path, route
	IsRemaining() bool
	Remaining() *uint8 // string, bool
}
