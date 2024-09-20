package chains

import (
	"github.com/steve-care-software/steve/domain/hash"
)

type transpile struct {
	bridge hash.Hash
	target hash.Hash
	next   Chain
}

func createTranspile(
	bridge hash.Hash,
	target hash.Hash,
) Transpile {
	return createTranspileInternally(bridge, target, nil)
}

func createTranspileWithNext(
	bridge hash.Hash,
	target hash.Hash,
	next Chain,
) Transpile {
	return createTranspileInternally(bridge, target, next)
}

func createTranspileInternally(
	bridge hash.Hash,
	target hash.Hash,
	next Chain,
) Transpile {
	out := transpile{
		bridge: bridge,
		target: target,
		next:   next,
	}

	return &out
}

// Bridge returns the bridge
func (obj *transpile) Bridge() hash.Hash {
	return obj.bridge
}

// Target returns the target
func (obj *transpile) Target() hash.Hash {
	return obj.target
}

// HasNext returns true if there is a next, false otherwise
func (obj *transpile) HasNext() bool {
	return obj.next != nil
}

// Next returns the next chain, if any
func (obj *transpile) Next() Chain {
	return obj.next
}
