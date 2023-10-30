package blockchains

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/identifiers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/links"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/pointers"
	"github.com/steve-care-software/steve/domain/hash"
)

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier hash.Hash) Builder
	WithHead(head blocks.Block) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Identifier() hash.Hash
	Head() blocks.Block
	Pointer() pointers.Pointer
	List() []resources.Resource
	ListByContainer(container []string) (resources.Resources, error)
	ListByContainerAndKind(container []string, kind uint) (resources.Resources, error)
	ExistsByIdentifier(identifier identifiers.Identifier) error
	ExistsByLink(link links.Link) (resources.Resource, error)
	FetchByIdentifier(identifier identifiers.Identifier) (resources.Resource, error)
	FetchByLink(link links.Link) (resources.Resource, error)
}

// Repository represents a repository
type Repository interface {
	List() []hash.Hash
	Retrieve(identifier hash.Hash) (Blockchain, error)
}

// Service represents a service
type Service interface {
	Save(blockchain Blockchain) error
	Delete(identifier hash.Hash) error
}
