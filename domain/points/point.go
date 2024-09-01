package points

import (
	"github.com/steve-care-software/steve/domain/points/bridges"
)

type point struct {
	bridge bridges.Bridge
	from   []byte
}

func createPoint(
	bridge bridges.Bridge,
	from []byte,
) Point {
	return createPointInternally(bridge, from)
}

func createPointInternally(
	bridge bridges.Bridge,
	from []byte,
) Point {
	out := point{
		bridge: bridge,
		from:   from,
	}

	return &out
}

// Bridge returns the bridge, if any
func (obj *point) Bridge() bridges.Bridge {
	return obj.bridge
}

// From returns the from data
func (obj *point) From() []byte {
	return obj.from
}
