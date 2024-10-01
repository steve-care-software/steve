package links

import "github.com/steve-care-software/steve/commons/hash"

type links struct {
	hash hash.Hash
	list []Link
}

func createLinks(
	hash hash.Hash,
	list []Link,
) Links {
	out := links{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *links) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *links) List() []Link {
	return obj.list
}
