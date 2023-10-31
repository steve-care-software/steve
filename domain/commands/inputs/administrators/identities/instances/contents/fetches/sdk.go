package fetches

import "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/fetches/properties"

// Builder represents a fetch builder
type Builder interface {
	Create() Builder
	WithAssignTo(assignTo string) Builder
	WithObject(object string) Builder
	WithProperty(property properties.Property) Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	AssignTo() string
	Object() string
	Property() properties.Property
}
