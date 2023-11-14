package results

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/kinds"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/results/resources"
)

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithKind(kind kinds.Kind) Builder
	WithBytes(bytes []byte) Builder
	WithPrevious(previous Result) Builder
	Now() (Result, error)
}

// Result represents the layer result
type Result interface {
	Hash() hash.Hash
	Resource() resources.Resource
	Kind() kinds.Kind
	Bytes() []byte
	HasPrevious() bool
	Previous() Result
}
