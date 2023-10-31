package instances

import "github.com/steve-care-software/steve/domain/commands/inputs/administrators/identities/instances/contents"

// Builder represents an instance builder
type Builder interface {
	Create() Builder
	WithObject(object string) Builder
	WithContent(content contents.Content) Builder
	Now() (Instance, error)
}

// Instance represents an instance command
type Instance interface {
	Object() string
	Content() contents.Content
}
