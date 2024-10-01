package points

import "github.com/steve-care-software/steve/engine/domain/hash"

type points struct {
	hash hash.Hash
	list []Point
}

func createPoints(
	hash hash.Hash,
	list []Point,
) Points {
	out := points{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *points) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *points) List() []Point {
	return obj.list
}
