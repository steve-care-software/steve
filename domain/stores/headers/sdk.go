package headers

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the header adapter
type Adapter interface {
	ToBytes(ins Header) ([]byte, error)
	ToInstance(data []byte) (Header, error)
}

// Builder represents the header builder
type Builder interface {
	Create() Builder
	WithRoot(root resources.Resources) Builder
	WithActivity(activity activities.Activity) Builder
	Now() (Header, error)
}

// Header represents the header
type Header interface {
	Hash() hash.Hash
	Root() resources.Resources
	HasActivity() bool
	Activity() activities.Activity
}
