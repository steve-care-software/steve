package points

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections"
)

// Points represents points
type Points interface {
	List() []Point
}

// Point represents a point
type Point interface {
	Connection() connections.Connection
	From() string
}
