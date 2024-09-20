package points

import "github.com/steve-care-software/steve/domain/pipelines/schemas/points/kinds"

const (
	// StructureList represents the list structure
	StructureList (uint8) = iota

	// StructureSet represents the set structure
	StructureSet

	// StructureSortedSet represents the sorted set structure
	StructureSortedSet
)

// Points represents points
type Points interface {
	List() []Point
}

// Point represents a point
type Point interface {
	Name() string
	Kind() kinds.Kind
	HasStructure() bool
	Structure() *uint8
}
