package points

import (
	"github.com/steve-care-software/steve/domain/points/bridges"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointBuilder creates a new point builder
func NewPointBuilder() PointBuilder {
	return createPointBuilder()
}

// Builder represents the points builder
type Builder interface {
	Create() Builder
	WithList(list []Point) Builder
	Now() (Points, error)
}

// Points represents points
type Points interface {
	List() []Point
}

// PointBuilder represents the point builder
type PointBuilder interface {
	Create() PointBuilder
	WithBridge(bridge bridges.Bridge) PointBuilder
	From(from []byte) PointBuilder
	Now() (Point, error)
}

// Point represents a point
type Point interface {
	Bridge() bridges.Bridge
	From() []byte
}
