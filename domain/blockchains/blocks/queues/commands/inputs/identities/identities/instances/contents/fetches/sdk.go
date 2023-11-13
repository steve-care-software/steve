package fetches

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities/identities/instances/contents/fetches/properties"

// Builder represents a fetch builder
type Builder interface {
	Create() Builder
	WithAssignTo(assignTo string) Builder
	WithProperty(property properties.Property) Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	AssignTo() string
	Property() properties.Property
}
