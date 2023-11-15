package resources

import (
	"github.com/steve-care-software/steve/vms/domain/resources/blocks"
	"github.com/steve-care-software/steve/vms/domain/resources/roots"
	"github.com/steve-care-software/steve/vms/libraries/hash"
)

// ActionFn represents an action func
type ActionFn = func() error

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithRoot(root roots.Root) Builder
	WithHead(head blocks.Block) Builder
	Now() (Resource, error)
}

// Resource represents a blockchain resource
type Resource interface {
	Hash() hash.Hash
	Root() roots.Root
	HasHead() bool
	Head() blocks.Block
}

// Repository represents a blockchain repository
type Repository interface {
	Retrieve(path string) (Resource, error)
}

// Service represents a blockchain service
type Service interface {
	Insert(path string, ins Resource) error
	Chain(original Resource, newBlock blocks.Block, action ActionFn) error
	Shrink(blockchain Resource, action ActionFn) error
}
