package inputs

import (
	"github.com/steve-care-software/steve/domain/commands/visitors/inputs/accounts"
	"github.com/steve-care-software/steve/domain/commands/visitors/inputs/administrators"
)

// Adapter represents a visitor's adapter
type Adapter interface {
	ToInput(bytes []byte) (Input, error)
	ToBytes(ins Input) ([]byte, error)
}

// Input represents visitor's input ommand
type Input interface {
	IsAccount() bool
	Account() accounts.Account
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
