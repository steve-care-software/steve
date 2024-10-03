package schemas

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
)

type schema struct {
	header      headers.Header
	points      []string
	connections connections.Connections
}

func createSchema(
	header headers.Header,
	points []string,
	connections connections.Connections,
) Schema {
	out := schema{
		header:      header,
		points:      points,
		connections: connections,
	}

	return &out
}

// Header returns the header
func (obj *schema) Header() headers.Header {
	return obj.header
}

// Points returns the points
func (obj *schema) Points() []string {
	return obj.points
}

// Connections returns the connections
func (obj *schema) Connections() connections.Connections {
	return obj.connections
}
