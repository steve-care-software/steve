package instances

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/administrators/instances/contents"

// Builder represents an instance builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithContent(content contents.Content) Builder
	Now() (Instance, error)
}

// Instance represents an instance command
type Instance interface {
	Name() string
	Content() contents.Content
}
