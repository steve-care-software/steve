package containers

import (
	"github.com/steve-care-software/steve/engine/domain/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
)

const (
	// FlagSingle represents the flag single
	FlagSingle (uint8) = iota

	// FlagMap represents the flag map
	FlagMap

	// FlagList represents the flag list
	FlagList

	// FlagSet represents the flag set
	FlagSet

	// FlagSortedSet represents the flag sorted_set
	FlagSortedSet
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() ContainerBuilder {
	hashAdapter := hash.NewAdapter()
	return createContainerBuilder(
		hashAdapter,
	)
}

// Builder represents the containers builder
type Builder interface {
	Create() Builder
	WithList(list []Container) Builder
	Now() (Containers, error)
}

// Containers represents containers
type Containers interface {
	Hash() hash.Hash
	List() []Container
}

// ContainerBuilder represents the container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithFlag(flag uint8) ContainerBuilder
	WithKind(kind kinds.Kind) ContainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Hash() hash.Hash
	Flag() uint8 // single, vector, list, set, sorted_set
	Kind() kinds.Kind
}
