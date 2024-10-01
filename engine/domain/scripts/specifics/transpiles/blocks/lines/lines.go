package lines

import "github.com/steve-care-software/steve/commons/hash"

type lines struct {
	hash hash.Hash
	list []Line
}

func createLines(
	hash hash.Hash,
	list []Line,
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
func (obj *lines) List() []Line {
	return obj.list
}
