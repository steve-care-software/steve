package access

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/writes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the access builder
type Builder interface {
	Create() Builder
	WithWrite(write writes.Write) Builder
	WithRead(read permissions.Permission) Builder
	Now() (Access, error)
}

// Access represents the access
type Access interface {
	Write() writes.Write
	HasRead() bool
	Read() permissions.Permission
}
