package resources

import "github.com/steve-care-software/steve/domain/stores/resources/pointers"

// Adapter represents a resource adapter
type Adapter interface {
	InstancesToBytes(ins Resources) ([]byte, error)
	BytesToInstances(data []byte) (Resources, error)
	InstanceToBytes(ins Resource) ([]byte, error)
	BytesToInstance(data []byte) (Resource, error)
}

// Resources represents resources
type Resources interface {
	List() []Resources
}

// Resource represents a resource
type Resource interface {
	Identifier() string
	Pointers() pointers.Pointers
}
