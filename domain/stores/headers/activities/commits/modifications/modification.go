package modifications

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
)

type modification struct {
	hash   hash.Hash
	insert resources.Resource
	save   resources.Resource
	delete string
}

func createModificationWithInsert(
	hash hash.Hash,
	insert resources.Resource,
) Modification {
	return createModificationInternally(hash, insert, nil, "")
}

func createModificationWithSave(
	hash hash.Hash,
	save resources.Resource,
) Modification {
	return createModificationInternally(hash, nil, save, "")
}

func createModificationWithDelete(
	hash hash.Hash,
	delete string,
) Modification {
	return createModificationInternally(hash, nil, nil, delete)
}

func createModificationInternally(
	hash hash.Hash,
	insert resources.Resource,
	save resources.Resource,
	delete string,
) Modification {
	out := modification{
		hash:   hash,
		insert: insert,
		save:   save,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *modification) Hash() hash.Hash {
	return obj.hash
}

// Identifier returns the identifier
func (obj *modification) Identifier() string {
	if obj.IsInsert() {
		return obj.insert.Identifier()
	}

	if obj.IsSave() {
		return obj.save.Identifier()
	}

	return obj.delete
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *modification) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *modification) Insert() resources.Resource {
	return obj.insert
}

// IsSave returns true if there is a save, false otherwise
func (obj *modification) IsSave() bool {
	return obj.save != nil
}

// Save returns the save, if any
func (obj *modification) Save() resources.Resource {
	return obj.save
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *modification) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *modification) Delete() string {
	return obj.delete
}
