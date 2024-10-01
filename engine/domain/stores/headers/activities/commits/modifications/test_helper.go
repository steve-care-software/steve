package modifications

import "github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources"

// NewModificationsForTests creates a new modifications for tests
func NewModificationsForTests(list []Modification) Modifications {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithInsertForTests creates a new modification by insert
func NewModificationWithInsertForTests(insert resources.Resource) Modification {
	ins, err := NewModificationBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithSaveForTests creates a new modification by save
func NewModificationWithSaveForTests(save resources.Resource) Modification {
	ins, err := NewModificationBuilder().Create().WithSave(save).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithDeleteForTests creates a new modification by delete
func NewModificationWithDeleteForTests(delete string) Modification {
	ins, err := NewModificationBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
