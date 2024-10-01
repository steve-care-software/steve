package headers

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	resourcesAdapter := resources.NewAdapter()
	activitiesAdapter := activities.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		resourcesAdapter,
		activitiesAdapter,
		builder,
	)
}

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
	ToInstance(data []byte) (Header, []byte, error)
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
	Map() (map[string]pointers.Pointer, error)
	NextPointerIndex() (*uint, error)
}
