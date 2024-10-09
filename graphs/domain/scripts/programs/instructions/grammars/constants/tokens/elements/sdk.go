package elements

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/constants/tokens/elements/references"
	"github.com/steve-care-software/steve/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithReference(reference references.Reference) Builder
	WithRule(rule string) Builder
	WithConstant(constant string) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	IsReference() bool
	Reference() references.Reference
	IsRule() bool
	Rule() string
	IsConstant() bool
	Constant() string
}
