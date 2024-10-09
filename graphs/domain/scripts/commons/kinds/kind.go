package kinds

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds/primitives"

// kind represents the implementation of the Kind interface
type kind struct {
	container Container
	pEngine   *uint8
	primitive primitives.Primitive
	isMap     bool
}

func createKindWithContainer(
	container Container,
) Kind {
	return createKindInternally(
		container,
		nil,
		nil,
		false,
	)
}

func createKindWithEngine(
	pEngine *uint8,
) Kind {
	return createKindInternally(
		nil,
		pEngine,
		nil,
		false,
	)
}

func createKindWithPrimitive(
	primitive primitives.Primitive,
) Kind {
	return createKindInternally(
		nil,
		nil,
		primitive,
		false,
	)
}

func createKindWithMap() Kind {
	return createKindInternally(
		nil,
		nil,
		nil,
		true,
	)
}

func createKindInternally(
	container Container,
	pEngine *uint8,
	primitive primitives.Primitive,
	isMap bool,
) Kind {
	return &kind{
		container: container,
		pEngine:   pEngine,
		primitive: primitive,
		isMap:     isMap,
	}
}

// IsContainer returns true if it is a container, false otherwise
func (obj *kind) IsContainer() bool {
	return obj.container != nil
}

// Container returns the container
func (obj *kind) Container() Container {
	return obj.container
}

// IsEngine returns true if it is an pEngine, false otherwise
func (obj *kind) IsEngine() bool {
	return obj.pEngine != nil
}

// Engine returns the pEngine
func (obj *kind) Engine() *uint8 {
	return obj.pEngine
}

// IsPrimitive returns true if it is a primitive, false otherwise
func (obj *kind) IsPrimitive() bool {
	return obj.primitive != nil
}

// Primitive returns the primitive
func (obj *kind) Primitive() primitives.Primitive {
	return obj.primitive
}

// IsMap returns true if it is a map, false otherwise
func (obj *kind) IsMap() bool {
	return obj.isMap
}
