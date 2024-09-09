package contents

import (
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithModification(resource modifications.Modification) Builder
	WithData(data []byte) Builder
	Now() (Content, error)
}

// Content represents a content
type Content interface {
	Modification() modifications.Modification
	Data() []byte
}
