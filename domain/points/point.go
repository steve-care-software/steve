package points

import (
	"github.com/steve-care-software/steve/domain/points/bridges"
	"github.com/steve-care-software/steve/domain/points/contexts"
)

type point struct {
	context contexts.Context
	bridge  bridges.Bridge
	from    []byte
}

func createPoint(
	context contexts.Context,
	bridge bridges.Bridge,
	from []byte,
) Point {
	return createPointInternally(context, bridge, from)
}

func createPointInternally(
	context contexts.Context,
	bridge bridges.Bridge,
	from []byte,
) Point {
	out := point{
		context: context,
		bridge:  bridge,
		from:    from,
	}

	return &out
}

// Context returns the context, if any
func (obj *point) Context() contexts.Context {
	return obj.context
}

// Bridge returns the bridge, if any
func (obj *point) Bridge() bridges.Bridge {
	return obj.bridge
}

// From returns the from data
func (obj *point) From() []byte {
	return obj.from
}
