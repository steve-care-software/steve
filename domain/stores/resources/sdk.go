package resources

import "github.com/steve-care-software/steve/domain/stores/resources/pointers"

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	pointersAdapter := pointers.NewAdapter()
	builder := NewBuilder()
	resourceBuilder := NewResourceBuilder()
	return createAdapter(
		pointersAdapter,
		builder,
		resourceBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	return createResourceBuilder()
}

// Adapter represents a resource adapter
type Adapter interface {
	InstancesToBytes(ins Resources) ([]byte, error)
	BytesToInstances(data []byte) (Resources, []byte, error)
	InstanceToBytes(ins Resource) ([]byte, error)
	BytesToInstance(data []byte) (Resource, []byte, error)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents resources
type Resources interface {
	List() []Resource
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithIdentifier(identifier string) ResourceBuilder
	WithPointers(pointers pointers.Pointers) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Identifier() string
	Pointers() pointers.Pointers
}
