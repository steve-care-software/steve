package resources

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

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

// Map returns the pointers map
func (obj *resources) Map() map[string]pointers.Pointer {
	output := map[string]pointers.Pointer{}
	for _, oneResource := range obj.list {
		identifier := oneResource.Identifier()
		output[identifier] = oneResource.Pointer()
	}

	return output
}
