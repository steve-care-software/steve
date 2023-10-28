package resources

import (
	"time"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/resources/links"
)

const (
	// SymbolLayerConstantValue represents the symbol layer constantValue
	SymbolLayerConstantValue (uint8) = iota
)

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithLink(link links.Link) Builder
	WithKind(kind uint8) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Link() links.Link
	Kind() uint8
	CreatedOn() time.Time
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents the resource repository
type Repository interface {
	List() (hash.Hash, error)
	ListByContainer(container []string) ([]hash.Hash, error)
	ListByContainerAndKind(container []string, kind uint8) ([]hash.Hash, error)
	ExistsByHashAndKind(hash hash.Hash, kind uint8) error
	ExistsByLink(link links.Link) (Resource, error)
	RetrieveByHashAndKind(hash hash.Hash, kind uint8) (Resource, error)
	RetrieveByLink(link links.Link) (Resource, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a resource service
type Service interface {
	Insert(res Resource) error
	Delete(res Resource) error
}
