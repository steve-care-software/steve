package contexts

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents"
)

// Context represents a context
type Context interface {
	Name() string
	Version() uint
	Content() contents.Content
	HasParent() bool
	Parent() hash.Hash
}
