package assignables

import "github.com/steve-care-software/steve/domain/stacks/assignables/administrators"

// Assignable represents an assignable
type Assignable interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
