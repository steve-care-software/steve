package schemas

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas/connections"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/schemas/points"
)

type schema struct {
	hash        hash.Hash
	head        heads.Head
	points      points.Points
	connections connections.Connections
}

func createSchema(
	hash hash.Hash,
	head heads.Head,
	points points.Points,
	connections connections.Connections,
) Schema {
	out := schema{
		hash:        hash,
		head:        head,
		points:      points,
		connections: connections,
	}

	return &out
}

// Hash returns the hash
func (obj *schema) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *schema) Head() heads.Head {
	return obj.head
}

// Points returns the points
func (obj *schema) Points() points.Points {
	return obj.points
}

// Connections returns the connections
func (obj *schema) Connections() connections.Connections {
	return obj.connections
}
