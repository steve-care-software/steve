package administrators

import "github.com/steve-care-software/steve/domain/accounts/administrators"

// Builder represents an administrator's builder
type Builder interface {
	Create() Builder
	WithCreate(create administrators.Administrator) Builder
	Now() (Administrator, error)
}

// Administrator represents an administrator's
type Administrator interface {
	IsCreate() bool
	Create() administrators.Administrator
}
