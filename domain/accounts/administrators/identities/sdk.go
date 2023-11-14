package identities

import "github.com/steve-care-software/steve/domain/dashboards"

// Builder represents an identities builder
type Builder interface {
	Create() Builder
	WithList(list []Identity) Builder
	Now() (Identities, error)
}

// Identities represents identities
type Identities interface {
	List() []Identity
	Fetch(index uint) (Identity, error)
	Delete(index uint) (Identities, error)
	Amount() uint
	Exceeds(amount uint) bool
}

// IdentityBuilder represents the identity builder
type IdentityBuilder interface {
	Create() IdentityBuilder
	WithName(name string) IdentityBuilder
	WithContainer(container []string) IdentityBuilder
	WithDashboard(dashboard dashboards.Dashboard) IdentityBuilder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Container() []string
	Dashboard() dashboards.Dashboard
}
