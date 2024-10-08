package contexts

import (
	"github.com/steve-care-software/steve/hash"
)

type context struct {
	hash   hash.Hash
	name   string
	parent hash.Hash
}

func createContext(
	hash hash.Hash,
	name string,
) Context {
	return createContextInternally(hash, name, nil)
}

func createContextWithParent(
	hash hash.Hash,
	name string,
	parent hash.Hash,
) Context {
	return createContextInternally(hash, name, parent)
}

func createContextInternally(
	hash hash.Hash,
	name string,
	parent hash.Hash,
) Context {
	out := context{
		hash:   hash,
		name:   name,
		parent: parent,
	}

	return &out
}

// Hash returns the hash
func (obj *context) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *context) Name() string {
	return obj.name
}

// HasParent returns true if there is parent, false otherwise
func (obj *context) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *context) Parent() hash.Hash {
	return obj.parent
}
