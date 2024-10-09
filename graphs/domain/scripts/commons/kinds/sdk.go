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

const (
	// EngineSelector represents the engine selector
	EngineSelector (uint8) = iota

	// EngineAST represents the engine AST
	EngineAST

	// EngineRoute represents the engine route
	EngineRoute

	// EngineSelect represents the engine select
	EngineSelect

	// EngineInsert represents the engine insert
	EngineInsert

	// EngineUpdate represents the engine update
	EngineUpdate

	// EngineDelete represents the engine delete
	EngineDelete

	// EngineBridges represents the engine bridges
	EngineBridges
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

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
	Now() (Kind, error)
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
