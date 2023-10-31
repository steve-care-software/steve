package visitors

import "github.com/steve-care-software/steve/domain/stencils"

// Builder represents the visitor's builder
type Builder interface {
	Create() Builder
	WithStencil(stencil stencils.Stencil) Builder
	Now() (Visitor, error)
}

// Visitor represents the visitor account
type Visitor interface {
	Stencil() stencils.Stencil
}

// Repository represents the visitor's repository
type Repository interface {
	Exists() (bool, error)
	Retrieve() (Visitor, error)
}

// Service represents the visitor's service
type Service interface {
	Insert(ins Visitor) error
	Delete(ins Visitor) error
}
