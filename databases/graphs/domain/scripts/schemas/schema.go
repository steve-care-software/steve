package schemas

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections"
)

type schema struct {
	head        heads.Head
	points      []string
	connections connections.Connections
}

func createSchema(
	head heads.Head,
	points []string,
	connections connections.Connections,
) Schema {
	out := schema{
		head:        head,
		points:      points,
		connections: connections,
	}

	return &out
}

// Head returns the head
func (obj *schema) Head() heads.Head {
	return obj.head
}

// Points returns the points
func (obj *schema) Points() []string {
	return obj.points
}

// Connections returns the connections
func (obj *schema) Connections() connections.Connections {
	return obj.connections
}
