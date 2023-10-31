package visitors

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts"
	"github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators"
)

// Visitor represents visitor command
type Visitor interface {
	IsAccount() bool
	Account() accounts.Account
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
