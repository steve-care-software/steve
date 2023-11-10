package assignables

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames/assignables/administrators"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames/assignables/identities"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Assignable represents an assignable
type Assignable interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsIdentity() bool
	Identity() identities.Identity
}
