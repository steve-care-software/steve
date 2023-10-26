package administrators

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators"
)

// Application represents the identity application
type Application interface {
	Retrieve() (administrators.Administrator, error)
	Save(criteria SaveCriteria) error
}

// SaveCriteria represents a save criteria
type SaveCriteria interface {
	Administrator() administrators.Administrator
	Password() []byte
	HasNewPassword() bool
	NewPassword() []byte
}
