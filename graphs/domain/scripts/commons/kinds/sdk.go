package kinds

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds/primitives"

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

// Container represents a container
type Container interface {
	Flag() uint8
	Kind() Kind
}
