package contexts

import (
	"github.com/steve-care-software/steve/domain/scripts/contexts/compensations"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents"
	"github.com/steve-care-software/steve/domain/scripts/contexts/roles"
)

// Context represents a context
type Context interface {
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
