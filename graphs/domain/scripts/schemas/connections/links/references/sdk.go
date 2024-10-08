package references

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewReferenceBuilder creates a new reference builder
func NewReferenceBuilder() ReferenceBuilder {
	return createReferenceBuilder()
}

// Builder represents the references builder
type Builder interface {
	Create() Builder
	WithList(list []Reference) Builder
	Now() (References, error)
}

// References represents references
type References interface {
	List() []Reference
}

// ReferenceBuilder represents the reference builder
type ReferenceBuilder interface {
	Create() ReferenceBuilder
	WithInternal(internal string) ReferenceBuilder
	WithExternal(external externals.External) ReferenceBuilder
	Now() (Reference, error)
}

type Reference interface {
	IsInternal() bool
	Internal() string
	IsExternal() bool
	External() externals.External
}
