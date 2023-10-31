package inputs

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/administrators"
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/identities"
	"github.com/steve-care-software/steve/domain/commands/shares/inputs/dashboards"
)

// Adapter represents a visitor's adapter
type Adapter interface {
	ToInput(bytes []byte) (Input, error)
	ToBytes(ins Input) ([]byte, error)
}

// Builder represents an input builder
type Builder interface {
	Create() Builder
	WithAdministrator(admin administrators.Administrator) Builder
	WithIdentities(identities identities.Identities) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	Now() (Input, error)
}

// Input represents admin's input ommand
type Input interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
	IsIdentities() bool
	Identities() identities.Identities
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
}
