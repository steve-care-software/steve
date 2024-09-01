package points

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/points/bridges"
	"github.com/steve-care-software/steve/domain/points/contexts"
)

type point struct {
	hash    hash.Hash
	context contexts.Context
	bridge  bridges.Bridge
	from    []byte
}

func createPoint(
	hash hash.Hash,
	context contexts.Context,
	bridge bridges.Bridge,
	from []byte,
) Point {
	return createPointInternally(hash, context, bridge, from)
}

func createPointInternally(
	hash hash.Hash,
	context contexts.Context,
	bridge bridges.Bridge,
	from []byte,
) Point {
	out := point{
		hash:    hash,
		context: context,
		bridge:  bridge,
		from:    from,
	}

	return &out
}

// Hash returns the hash, if any
func (obj *point) Hash() hash.Hash {
	return obj.hash
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
