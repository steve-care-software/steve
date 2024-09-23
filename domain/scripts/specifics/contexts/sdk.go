package contexts

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/domain/scripts/specifics/contexts/contents"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

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
