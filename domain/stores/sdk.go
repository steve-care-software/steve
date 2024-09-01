package stores

import "github.com/steve-care-software/steve/domain/hash"

// ListAdapter represents a list adapter
type ListAdapter interface {
	ToBytes(list List) ([]byte, error)
	ToInstance(values []byte) (List, error)
}

// List represents a list
type List interface {
	Hash() hash.Hash
	Values() []string
}
