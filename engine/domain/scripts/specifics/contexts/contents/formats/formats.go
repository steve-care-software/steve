package formats

import "github.com/steve-care-software/steve/commons/hash"

type formats struct {
	hash hash.Hash
	list []Format
}

func createFormats(
	hash hash.Hash,
	list []Format,
) Formats {
	out := formats{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *formats) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *formats) List() []Format {
	return obj.list
}
