package successes

import "github.com/steve-care-software/steve/domain/accounts/administrators/identities"

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithAmount(amount uint) Builder
	WithAtIndex(atIndex identities.Identity) Builder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	IsAmount() bool
	Amount() *uint
	IsAtIndex() bool
	AtIndex() identities.Identity
}
