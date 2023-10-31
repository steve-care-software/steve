package administrators

import "github.com/steve-care-software/steve/domain/commands/visitors/inputs/administrators/creates"

// Administrator represents an administrator
type Administrator interface {
	IsCreate() bool
	Create() creates.Create
}
