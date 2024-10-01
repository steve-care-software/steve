package modifications

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type modifications struct {
	hash hash.Hash
	list []Modification
}

func createModifications(
	hash hash.Hash,
	list []Modification,
) Modifications {
	out := modifications{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *modifications) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *modifications) List() []Modification {
	return obj.list
}

// Fetch fetches a pointer
func (obj *modifications) Map() (map[string]pointers.Pointer, []string) {
	deleted := []string{}
	ptrMap := map[string]pointers.Pointer{}
	for _, oneModification := range obj.list {
		if oneModification.IsDelete() {
			deleted = append(deleted, oneModification.Delete())
			continue
		}

		resource := oneModification.Insert()
		if oneModification.IsSave() {
			resource = oneModification.Save()
		}

		identifier := resource.Identifier()
		ptrMap[identifier] = resource.Pointer()
	}

	return ptrMap, deleted
}
