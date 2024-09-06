package resources

import "github.com/steve-care-software/steve/domain/hash"

type resources struct {
	hash hash.Hash
	list []Resource
}

func createResources(
	hash hash.Hash,
	list []Resource,
) Resources {
	out := resources{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *resources) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of resources
func (obj *resources) List() []Resource {
	return obj.list
}
