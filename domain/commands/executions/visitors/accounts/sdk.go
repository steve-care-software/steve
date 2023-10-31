package accounts

import "github.com/steve-care-software/steve/domain/accounts/visitors"

// Builder represents an acocunt builder
type Builder interface {
	Create() Builder
	WithCreate(create visitors.Visitor) Builder
	Now() (Account, error)
}

// Account represents an account
type Account interface {
	IsCreate() bool
	Create() visitors.Visitor
}
