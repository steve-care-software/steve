package points

import "github.com/steve-care-software/steve/domain/relations/data/connections"

type point struct {
	connection connections.Connection
	from       []byte
}

func createPoint(
	connection connections.Connection,
	from []byte,
) Point {
	out := point{
		connection: connection,
		from:       from,
	}

	return &out
}

// Connection rreturns the connection
func (obj *point) Connection() connections.Connection {
	return obj.connection
}

// From returns the from data
func (obj *point) From() []byte {
	return obj.from
}
