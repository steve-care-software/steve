package actions

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/transactions"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithSigner(signer signers.Signer) Builder
	WithBlockchain(blockchain blockchains.Blockchain) Builder
	Now() (Application, error)
}

// Application represents an action blockchain application
type Application interface {
	Insert(trx transactions.Transaction) error
	Push() error
	Commit(message string) error
	Rollback() error
	Cancel()
}
