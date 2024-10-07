package writes

import "github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the write builder
type Builder interface {
	Create() Builder
	WithModify(modify permissions.Permissions) Builder
	WithReview(review permissions.Permissions) Builder
	Now() (Write, error)
}

// Write represents the write permissions
type Write interface {
	Modify() permissions.Permissions
	HasReview() bool
	Review() permissions.Permissions
}
