package assignables

import "github.com/steve-care-software/steve/domain/accounts/administrators"

// Assignable represents an assignable
type Assignable interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
