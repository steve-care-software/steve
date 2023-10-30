package resources

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/pointers"
	"github.com/steve-care-software/steve/domain/hash"
)

// Builder represents resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents resources
type Resources interface {
	List() []Resource
	Pointer() pointers.Pointer
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithHeader(header headers.Header) ResourceBuilder
	WithPointer(pointer pointers.Pointer) ResourceBuilder
	CreatedOn(createdOn time.Time) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Header() headers.Header
	Pointer() pointers.Pointer
	CreatedOn() time.Time
}
