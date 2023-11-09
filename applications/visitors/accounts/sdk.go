package accounts

import (
	executions "github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/visitors/accounts"
)

// Application represents the application
type Application interface {
	Execute(account inputs.Account) (executions.Account, error)
}
