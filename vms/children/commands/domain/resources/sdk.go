package resources

import "github.com/steve-care-software/steve/vms/libraries/hash"

// Builder represents a resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents a list of command resources
type Resources interface {
	List() []Resource
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a command resource
type Resource interface {
	Hash() hash.Hash
}
