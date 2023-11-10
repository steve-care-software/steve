package parameters

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/parameters/kinds"
)

type parameter struct {
	hash hash.Hash
	name string
	kind kinds.Kind
}

func createParameter(
	hash hash.Hash,
	name string,
	kind kinds.Kind,
) Parameter {
	out := parameter{
		hash: hash,
		name: name,
		kind: kind,
	}

	return &out
}

// Hash returns the hash
func (obj *parameter) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *parameter) Name() string {
	return obj.name
}

// Kind returns the kind
func (obj *parameter) Kind() kinds.Kind {
	return obj.kind
}
