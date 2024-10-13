package params

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the params builder
type Builder interface {
	Create() Builder
	WithKind(kind kinds.Kind) Builder
	WithInternal(internal string) Builder
	WithExternal(external string) Builder
	IsMandatory() Builder
	Now() (Params, error)
}

// Params represents params
type Params interface {
	Kind() kinds.Kind
	Internal() string
	External() string
	IsMandatory() bool
}
