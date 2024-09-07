package modifications

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/commits/modifications/resources"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewModificationBuilder creates a new modification builder
func NewModificationBuilder() ModificationBuilder {
	hashAdapter := hash.NewAdapter()
	return createModificationBuilder(
		hashAdapter,
	)
}

// Adapter represents the modifications adapter
type Adapter interface {
	InstancesToBytes(ins Modifications) ([]byte, error)
	BytesToInstances(data []byte) (Modifications, error)
	InstanceToBytes(ins Modification) ([]byte, error)
	BytesToInstance(data []byte) (Modification, error)
}

// Builder represents the modifications builder
type Builder interface {
	Create() Builder
	WithList(list []Modification) Builder
	Now() (Modifications, error)
}

// Modifications represents modifications
type Modifications interface {
	Hash() hash.Hash
	List() []Modification
}

// ModificationBuilder represents the modification builder
type ModificationBuilder interface {
	Create() ModificationBuilder
	WithInsert(insert resources.Resource) ModificationBuilder
	WithSave(save resources.Resource) ModificationBuilder
	WithDelete(delete string) ModificationBuilder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() resources.Resource
	IsSave() bool
	Save() resources.Resource
	IsDelete() bool
	Delete() string
}
