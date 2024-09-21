package points

import (
	"github.com/steve-care-software/steve/domain/hash"
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

// Points represents points
type Points interface {
	Hash() hash.Hash
	List() []Point
}

// Point represents a point
type Point interface {
	Hash() hash.Hash
	Name() string
	Kind() uint8
	HasStructure() bool
	Structure() *uint8
}
