package administrators

import (
	"github.com/steve-care-software/steve/domain/commands/executions/shares/administrators/creates"
)

// Builder represents an administrator's builder
type Builder interface {
	Create() Builder
	WithCreate(create creates.Create) Builder
	Now() (Administrator, error)
}

// Administrator represents an administrator's
type Administrator interface {
	IsCreate() bool
	Create() creates.Create
}
