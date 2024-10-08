package bridges

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/bridges/connections"
	"github.com/steve-care-software/steve/hash"
)

type bridge struct {
	hash        hash.Hash
	head        heads.Head
	origin      string
	target      string
	connections connections.Connections
}

func createBridge(
	hash hash.Hash,
	head heads.Head,
	origin string,
	target string,
	connections connections.Connections,
) Bridge {
	out := bridge{
		hash:        hash,
		head:        head,
		origin:      origin,
		target:      target,
		connections: connections,
	}

	return &out
}

// Hash returns the hash
func (obj *bridge) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *bridge) Head() heads.Head {
	return obj.head
}

// Origin returns the origin
func (obj *bridge) Origin() string {
	return obj.origin
}

// Target returns the target
func (obj *bridge) Target() string {
	return obj.target
}

// Connections returns the connections
func (obj *bridge) Connections() connections.Connections {
	return obj.connections
}
