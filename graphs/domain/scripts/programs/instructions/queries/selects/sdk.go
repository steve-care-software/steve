package selects

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

// Select represents a select
type Select interface {
	IsDelete() bool
	Externals() externals.Externals
	HasCondition() bool
	Condition() conditions.Condition
}
