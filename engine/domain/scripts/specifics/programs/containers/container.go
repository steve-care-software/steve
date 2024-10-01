package containers

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"
)

type container struct {
	hash hash.Hash
	flag uint8
	kind kinds.Kind
}

func createContainer(
	hash hash.Hash,
	flag uint8,
	kind kinds.Kind,
) Container {
	out := container{
		hash: hash,
		flag: flag,
		kind: kind,
	}

	return &out
}

// Hash returns the hash
func (obj *container) Hash() hash.Hash {
	return obj.hash
}

// Flag returns the flag
func (obj *container) Flag() uint8 {
	return obj.flag
}

// Kind returns the kind
func (obj *container) Kind() kinds.Kind {
	return obj.kind
}
