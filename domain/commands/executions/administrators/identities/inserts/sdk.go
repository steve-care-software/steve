package inserts

import "github.com/steve-care-software/steve/domain/accounts/administrators/identities"

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithIdentities(identities identities.Identities) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Identities() identities.Identities
}
