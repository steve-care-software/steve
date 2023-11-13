package successes

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/shares/administrators/creates"

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithAdministrator(admin creates.Create) Builder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	IsAdministrator() bool
	Administrator() creates.Create
}
