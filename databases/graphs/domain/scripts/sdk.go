package scripts

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas"

// NewBuilder creates a new script builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the script adapter
type Adapter interface {
	ToScript(input []byte) (Script, []byte, error)
}

// Builder represents the script builder
type Builder interface {
	Create() Builder
	WithSchema(schema schemas.Schema) Builder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	IsSchema() bool
	Schema() schemas.Schema
}
