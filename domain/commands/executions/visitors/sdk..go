package visitors

import (
	"github.com/steve-care-software/steve/domain/commands/executions/visitors/accounts"
	"github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithAccount(account accounts.Account) Builder
	WithAdministrator(administrator administrators.Administrator) Builder
	Now() (Visitor, error)
}

// Visitor represents a visitor
type Visitor interface {
	IsAccount() bool
	Account() accounts.Account
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
