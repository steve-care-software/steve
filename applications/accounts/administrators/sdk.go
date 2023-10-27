package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
)

// Application represents the identity application
type Application interface {
	Retrieve() (administrators.Administrator, error)
	Save(criteria SaveCriteria) error
}

// SaveCriteriaBuilder represents a save criteria builder
type SaveCriteriaBuilder interface {
	Create() SaveCriteriaBuilder
	WithAdministrator(admin administrators.Administrator) SaveCriteriaBuilder
	WithPassword(password []byte) SaveCriteriaBuilder
	WithNewPassword(newPassword []byte) SaveCriteriaBuilder
	Now() (SaveCriteria, error)
}

// SaveCriteria represents a save criteria
type SaveCriteria interface {
	Administrator() administrators.Administrator
	Password() []byte
	HasNewPassword() bool
	NewPassword() []byte
}
