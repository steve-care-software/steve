package resources

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/dashboards/stencils/pointers"
)

// Builder represents the resource builder
type Builder interface {
	Create() Builder
	WithLayer(layer pointers.Pointer) Builder
	WithTrigger(trigger pointers.Pointer) Builder
	Now() (Resource, error)
}

// Resource represents the result resource
type Resource interface {
	Hash() hash.Hash
	Layer() pointers.Pointer
	HasTrigger() bool
	Trigger() pointers.Pointer
}
