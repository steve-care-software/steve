package values

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithUsername(username string) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithHasIdentities(hasIdentities bool) Builder
	WithIdentities(identities identities.Identities) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsUsername() bool
	Username() string
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsHasIdentities() bool
	HasIdentities() *bool
	IsIdentities() bool
	Identities() identities.Identities
}
