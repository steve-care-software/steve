package visitors

import (
	"github.com/steve-care-software/steve/domain/commands/executions/visitors/administrators"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithAdministrator(administrator administrators.Administrator) Builder
	Now() (Visitor, error)
}

// Visitor represents a visitor
type Visitor interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
