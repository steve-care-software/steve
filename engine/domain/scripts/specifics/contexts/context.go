package contexts

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/contexts/contents"
)

type context struct {
	hash    hash.Hash
	head    heads.Head
	content contents.Content
	parent  string
}

func createContextWithParent(
	hash hash.Hash,
	head heads.Head,
	content contents.Content,
	parent string,
) Context {
	return createContextInternally(
		hash,
		head,
		content,
		parent,
	)
}

func createContext(
	hash hash.Hash,
	head heads.Head,
	content contents.Content,
) Context {
	return createContextInternally(
		hash,
		head,
		content,
		"",
	)
}

func createContextInternally(
	hash hash.Hash,
	head heads.Head,
	content contents.Content,
	parent string,
) Context {
	out := context{
		hash:    hash,
		head:    head,
		content: content,
		parent:  parent,
	}

	return &out
}

// Hash returns the hash
func (obj *context) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *context) Head() heads.Head {
	return obj.head
}

// Content returns the content
func (obj *context) Content() contents.Content {
	return obj.content
}

// HasParent returns true if there is a parent, false otherwise
func (obj *context) HasParent() bool {
	return obj.parent != ""
}

// Parent returns the parent, if any
func (obj *context) Parent() string {
	return obj.parent
}
