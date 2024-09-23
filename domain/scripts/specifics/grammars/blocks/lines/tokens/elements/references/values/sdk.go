package values

import "github.com/steve-care-software/steve/domain/hash"

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithBlock(block string) Builder
	WithConstant(constant string) Builder
	WithRule(rule string) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	IsBlock() bool
	Block() string
	IsConstant() bool
	Constant() string
	IsRule() bool
	Rule() string
}
