package headers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
)

type header struct {
	hash     hash.Hash
	root     resources.Resources
	activity activities.Activity
}

func createHeader(
	hash hash.Hash,
	root resources.Resources,
) Header {
	return createHeaderInternally(hash, root, nil)
}

func createHeaderWithActivity(
	hash hash.Hash,
	root resources.Resources,
	activity activities.Activity,
) Header {
	return createHeaderInternally(hash, root, activity)
}

func createHeaderInternally(
	hash hash.Hash,
	root resources.Resources,
	activity activities.Activity,
) Header {
	out := header{
		hash:     hash,
		root:     root,
		activity: activity,
	}

	return &out
}

// Hash returns the hash
func (obj *header) Hash() hash.Hash {
	return obj.hash
}

// Root returns the root
func (obj *header) Root() resources.Resources {
	return obj.root
}

// HasActivity returns true if there is an activity, false otherwise
func (obj *header) HasActivity() bool {
	return obj.activity != nil
}

// Activity returns the activity, if any
func (obj *header) Activity() activities.Activity {
	return obj.activity
}
