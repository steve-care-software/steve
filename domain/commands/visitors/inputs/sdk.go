package inputs

import (
	"github.com/steve-care-software/steve/domain/commands/visitors/inputs/accounts"
	"github.com/steve-care-software/steve/domain/commands/visitors/inputs/administrators"
)

// Adapter represents a visitor's adapter
type Adapter interface {
	ToVisitor(bytes []byte) (Visitor, error)
	ToBytes(ins Visitor) ([]byte, error)
}

// Visitor represents visitor's commands
type Visitor interface {
	IsAccount() bool
	Account() accounts.Account
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
