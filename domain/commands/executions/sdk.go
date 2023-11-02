package executions

import (
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/administrators"
	"github.com/steve-care-software/steve/domain/commands/executions/visitors"
)

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithAdministrator(administrator administrators.Administrator) Builder
	WithVisitor(visitor visitors.Visitor) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
	IsVisitor() bool
	Visitor() visitors.Visitor
}
