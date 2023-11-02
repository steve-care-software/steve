package interpreters

import "github.com/steve-care-software/steve/domain/assignables"

// Application represents the interpreter application
type Application interface {
	Retrieve(name string) (assignables.Assignable, error)
}
