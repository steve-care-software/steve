package visitors

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/administrators/creates"
)

// Builder represents a visitor builder
type Builder interface {
	Create() Builder
	WithAdministrator(admin creates.Create) Builder
	Now() (Visitor, error)
}

// Visitor represents visitor command
type Visitor interface {
	IsAdministrator() bool
	Administrator() creates.Create
}
