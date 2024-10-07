package scripts

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access"

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
