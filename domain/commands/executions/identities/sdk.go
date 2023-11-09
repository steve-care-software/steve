package identities

import (
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities"
	"github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	IsIdentity() bool
	Identity() identities.Identity
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
}
