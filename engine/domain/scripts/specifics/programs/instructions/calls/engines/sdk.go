package engines

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/calls/functions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the engine builder
type Builder interface {
	Create() Builder
	WithScope(scope uint8) Builder
	WithFunction(function functions.Function) Builder
	Now() (Engine, error)
}

// Engine represents an engine call
type Engine interface {
	Hash() hash.Hash
	Scope() uint8 // role, identity, etc
	Function() functions.Function
}