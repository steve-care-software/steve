package points

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections"
)

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
	WithConnection(connection connections.Connection) PointBuilder
	From(from string) PointBuilder
	Now() (Point, error)
}

// Point represents a point
type Point interface {
	Connection() connections.Connection
	From() string
}
