package access

import (
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/writes"
)

// Builder represents the access builder
type Builder interface {
	Create() Builder
	WithWrite(write writes.Write) Builder
	WithRead(read permissions.Permissions) Builder
	Now() (Access, error)
}

// Access represents the access
type Access interface {
	Write() writes.Write
	HasRead() bool
	Read() permissions.Permissions
}
