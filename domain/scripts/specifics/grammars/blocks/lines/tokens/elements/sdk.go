package elements

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens/elements/references"
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
	WithBlock(block string) Builder
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
	IsBlock() bool
	Block() string
}
