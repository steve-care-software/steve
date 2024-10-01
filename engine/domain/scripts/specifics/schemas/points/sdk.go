package points

import (
	"github.com/steve-care-software/steve/commons/hash"
)

const (
	// StructureList represents the list structure
	StructureList (uint8) = iota

	// StructureSet represents the set structure
	StructureSet

	// StructureSortedSet represents the sorted set structure
	StructureSortedSet

	// StructureVector represents the vector structure
	StructureVector
)

const (
	// KindBytes represents the bytes kind
	KindBytes (uint8) = iota

	// KindInt represents the int kind
	KindInt

	// KindUint represents the uint kind
	KindUint

	// KindFloat represents the float kind
	KindFloat
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewPointBuilder creates a new point builder
func NewPointBuilder() PointBuilder {
	hashAdapter := hash.NewAdapter()
	return createPointBuilder(
		hashAdapter,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Point) Builder
	Now() (Points, error)
}

// Points represents points
type Points interface {
	Hash() hash.Hash
	List() []Point
}

// PointBuilder represents the point builder
type PointBuilder interface {
	Create() PointBuilder
	WithName(name string) PointBuilder
	WithKind(kind uint8) PointBuilder
	WithStructure(structure uint8) PointBuilder
	Now() (Point, error)
}

// Point represents a point
type Point interface {
	Hash() hash.Hash
	Name() string
	Kind() uint8
	HasStructure() bool
	Structure() *uint8
}
