package contexts

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/compensations"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents"
	"github.com/steve-care-software/steve/domain/scripts/contexts/roles"
)

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version uint) Builder
	WithContent(content contents.Content) Builder
	WithParent(parent string) Builder
	WithRole(role roles.Role) Builder
	WithCompensation(compensation compensations.Compensation) Builder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Hash() hash.Hash
	Name() string
	Version() uint
	Content() contents.Content
	HasParent() bool
	Parent() string
	HasRole() bool
	Role() roles.Role
	HasCompensation() bool
	Compensation() compensations.Compensation
}
