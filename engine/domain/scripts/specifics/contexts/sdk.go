package contexts

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents"
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
	WithHead(head heads.Head) Builder
	WithContent(content contents.Content) Builder
	WithParent(parent string) Builder
	Now() (Context, error)
}

// Context represents a context
type Context interface {
	Hash() hash.Hash
	Head() heads.Head
	Content() contents.Content
	HasParent() bool
	Parent() string
}
