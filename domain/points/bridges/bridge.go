package bridges

import (
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/hash"
)

type bridge struct {
	hash       hash.Hash
	connection connections.Connection
	weight     float32
}

func createBridge(
	hash hash.Hash,
	connection connections.Connection,
) Bridge {
	return createBridgeInternally(hash, connection, 0.0)
}

func createBridgeWithWeight(
	hash hash.Hash,
	connection connections.Connection,
	weight float32,
) Bridge {
	return createBridgeInternally(hash, connection, weight)
}

func createBridgeInternally(
	hash hash.Hash,
	connection connections.Connection,
	weight float32,
) Bridge {
	out := bridge{
		hash:       hash,
		connection: connection,
		weight:     weight,
	}

	return &out
}

// Hash returns the hash
func (obj *bridge) Hash() hash.Hash {
	return obj.hash
}

// Connection returns the connection
func (obj *bridge) Connection() connections.Connection {
	return obj.connection
}

// HasWeight returns true if there is a weight, false otherwise
func (obj *bridge) HasWeight() bool {
	return obj.weight > 0.0
}

// Weight returns the weight
func (obj *bridge) Weight() float32 {
	return obj.weight
}
