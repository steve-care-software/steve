package reverses

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/parsers/domain/grammars/blocks/lines/tokens/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the reverse builder
type Builder interface {
	Create() Builder
	WithEscape(escape elements.Element) Builder
	Now() (Reverse, error)
}

// Reverse represents the reverse builder
type Reverse interface {
	Hash() hash.Hash
	HasEscape() bool
	Escape() elements.Element
}
