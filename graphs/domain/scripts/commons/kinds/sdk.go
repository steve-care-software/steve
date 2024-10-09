package kinds

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds/primitives"

const (
	// ContainerList represents a list
	ContainerList (uint8) = iota

	// ContainerSet represents a set
	ContainerSet

	// ContainerSortedSet represents a sorted set
	ContainerSortedSet
)

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() ContainerBuilder {
	return createContainerBuilder()
}

// Builder represents the kind builder
type Builder interface {
	Create() Builder
	WithContainer(container Container) Builder
	WithEngine(engine uint8) Builder
	WithPrimitive(primitive primitives.Primitive) Builder
	IsMap() Builder
	Now() Kind
}

// Kind represents a kind
type Kind interface {
	IsContainer() bool
	Container() Container
	IsEngine() bool
	Engine() *uint8
	IsPrimitive() bool
	Primitive() primitives.Primitive
	IsMap() bool
}

// ContainerBuilder represents a container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithFlag(flag uint8) ContainerBuilder
	WithKind(kind Kind) ContainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Flag() uint8
	Kind() Kind
}
