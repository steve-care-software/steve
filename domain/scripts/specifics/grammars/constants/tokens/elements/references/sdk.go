package references

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/constants/tokens/elements/references/values"
)

// Builder represents the reference builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar string) Builder
	WithValue(value values.Value) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Hash() hash.Hash
	Grammar() string
	Value() values.Value
}