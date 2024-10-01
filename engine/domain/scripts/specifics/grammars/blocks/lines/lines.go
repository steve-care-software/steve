package lines

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/grammars/blocks/lines/tokens"
)

type lines struct {
	hash hash.Hash
	list []tokens.Tokens
}

func createLines(
	hash hash.Hash,
	list []tokens.Tokens,
) Lines {
	out := lines{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *lines) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *lines) List() []tokens.Tokens {
	return obj.list
}
