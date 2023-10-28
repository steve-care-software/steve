package returns

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/returns/kinds"
)

type ret struct {
	hash   hash.Hash
	output []byte
	kind   kinds.Kind
}

func createReturn(
	hash hash.Hash,
	output []byte,
	kind kinds.Kind,
) Return {
	out := ret{
		hash:   hash,
		output: output,
		kind:   kind,
	}

	return &out
}

// Hash returns the hash
func (obj *ret) Hash() hash.Hash {
	return obj.hash
}

// Output returns the output
func (obj *ret) Output() []byte {
	return obj.output
}

// Kind returns the kind
func (obj *ret) Kind() kinds.Kind {
	return obj.kind
}
