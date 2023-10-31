package accounts

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts/creates"
)

// Account represents a visitor's account
type Account interface {
	IsCreate() bool
	Create() creates.Create
}
