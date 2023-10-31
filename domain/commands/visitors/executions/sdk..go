package executions

import (
	"github.com/steve-care-software/steve/domain/commands/visitors/executions/accounts"
	"github.com/steve-care-software/steve/domain/commands/visitors/executions/administrators"
)

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	WithAdministrator(administrator administrators.Administrator) Builder
	Now() (Execution, error)
}

// Execution represents an executions
type Execution interface {
	IsAccount() bool
	Account() accounts.Account
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
