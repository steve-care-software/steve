package connections

import "github.com/steve-care-software/steve/commons/hash"

type connections struct {
	hash hash.Hash
	list []Connection
}

func createConnections(
	hash hash.Hash,
	list []Connection,
) Connections {
	out := connections{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *connections) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *connections) List() []Connection {
	return obj.list
}
