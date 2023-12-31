package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/accounts/credentials"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Builder represents an administrator's builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithIdentities(identities identities.Identities) Builder
	Now() (Administrator, error)
}

// Administrator represents an administrator
type Administrator interface {
	Username() string
	Dashboard() dashboards.Dashboard
	HasIdentities() bool
	Identities() identities.Identities
}

// Repository represents an administrator's repository
type Repository interface {
	IsInitialized() (bool, error)
	List() ([]string, error)
	Exists(username string) (bool, error)
	Retrieve(credentials credentials.Credentials) (Administrator, error)
}

// Service represents an administrator's service
type Service interface {
	Insert(admin Administrator, password []byte) error
	Save(admin Administrator, password []byte, newPassword []byte) error
	Delete(credentials credentials.Credentials) error
}
