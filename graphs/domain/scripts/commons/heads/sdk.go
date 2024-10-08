package heads

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the head builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	WithAccess(access access.Access) Builder
	Now() (Head, error)
}

// Head represents the head
type Head interface {
	Name() string
	Version() uint
	Access() access.Access
}
