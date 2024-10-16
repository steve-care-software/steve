package lines

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/steve/hash"
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
