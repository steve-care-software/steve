package administrators

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/shares/dashboards"
)

// Builder represents an administrator builder
type Builder interface {
	Create() Builder
	WithAdministrator(admin administrators.Administrator) Builder
	WithIdentities(identities identities.Identities) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	Now() (Administrator, error)
}

// Administrator represents administrator
type Administrator interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
	IsIdentities() bool
	Identities() identities.Identities
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
}
