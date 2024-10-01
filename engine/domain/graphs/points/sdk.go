package points

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/graphs/points/bridges"
	"github.com/steve-care-software/steve/engine/domain/graphs/points/contexts"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointBuilder creates a new point builder
func NewPointBuilder() PointBuilder {
	hashAdapter := hash.NewAdapter()
	return createPointBuilder(
		hashAdapter,
	)
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
	WithContext(context contexts.Context) PointBuilder
	WithBridge(bridge bridges.Bridge) PointBuilder
	From(from []byte) PointBuilder
	Now() (Point, error)
}

// Point represents a point
type Point interface {
	Hash() hash.Hash
	Context() contexts.Context
	Bridge() bridges.Bridge
	From() []byte
}
