package writes

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the write builder
type Builder interface {
	Create() Builder
	WithModify(modify permissions.Permission) Builder
	WithReview(review permissions.Permission) Builder
	Now() (Write, error)
}

// Write represents the write permissions
type Write interface {
	Modify() permissions.Permission
	HasReview() bool
	Review() permissions.Permission
}
