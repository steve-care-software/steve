package fetches

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/fetches/values"

// Builder represents a fetch builder
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
