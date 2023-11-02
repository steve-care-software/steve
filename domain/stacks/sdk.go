package stacks

import "github.com/steve-care-software/steve/domain/stacks/assignables"

// Stack represents the stack
type Stack interface {
	Fetch(name string) (assignables.Assignable, error)
}
