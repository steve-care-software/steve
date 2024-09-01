package points

import (
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/points/bridges"
)

type point struct {
	connection connections.Connection
	from       []byte
	bridge     bridges.Bridge
}

func createPoint(
	connection connections.Connection,
	from []byte,
) Point {
	return createPointInternally(connection, from, nil)
}

func createPointWithBridge(
	connection connections.Connection,
	from []byte,
	bridge bridges.Bridge,
) Point {
	return createPointInternally(connection, from, bridge)
}

func createPointInternally(
	connection connections.Connection,
	from []byte,
	bridge bridges.Bridge,
) Point {
	out := point{
		connection: connection,
		from:       from,
		bridge:     bridge,
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

// HasBridge returns true if there is a bridge, false otherwise
func (obj *point) HasBridge() bool {
	return obj.bridge != nil
}

// Bridge returns the bridge, if any
func (obj *point) Bridge() bridges.Bridge {
	return obj.bridge
}
