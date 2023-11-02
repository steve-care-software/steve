package fetches

import (
	"github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards/successes/fetches/values"
)

// Builder represents the fetch builder
type Builder interface {
	Create() Builder
	WithAssignTo(assignTo string) Builder
	WithValue(value values.Value) Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	AssignTo() string
	Value() values.Value
}
