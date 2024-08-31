package relations

import (
	"github.com/steve-care-software/steve/domain/relations/data/connections"
	"github.com/steve-care-software/steve/domain/relations/data/points"
)

// Relation represents a relation
type Relation interface {
	Connections() connections.Connections
	HasPoints() bool
	Points() points.Points
}
