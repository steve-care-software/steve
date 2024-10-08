package headers

import (
	"github.com/steve-care-software/steve/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources"
)

// NewHeaderForTests creates a new header for tests
func NewHeaderForTests(root resources.Resources) Header {
	ins, err := NewBuilder().Create().WithRoot(root).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewHeaderForTests creates a new header for tests
func NewHeaderWithActivityForTests(root resources.Resources, activity activities.Activity) Header {
	ins, err := NewBuilder().Create().WithRoot(root).WithActivity(activity).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
