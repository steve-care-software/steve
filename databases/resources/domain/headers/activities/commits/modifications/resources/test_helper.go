package resources

import "github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"

// NewResourcesForTests creates a new resources for tests
func NewResourcesForTests(list []Resource) Resources {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceForTests creates a new resource for tests
func NewResourceForTests(identifier string, pointer pointers.Pointer) Resource {
	ins, err := NewResourceBuilder().Create().
		WithIdentifier(identifier).
		WithPointer(pointer).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
