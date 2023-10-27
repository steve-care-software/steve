package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Administrator represents an administrator
type Administrator interface {
	Username() string
	Dashboard() dashboards.Dashboard
	HasIdentities() bool
	Identities() identities.Identities
}

// Repository represents an administrator's repository
type Repository interface {
	Retrieve(username string, password []byte) (Administrator, error)
}

// Service represents an administrator's service
type Service interface {
	Insert(admin Administrator, password []byte, newPassword []byte) error
	Delete(username string, password []byte) error
}
