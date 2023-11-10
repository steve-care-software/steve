package identities

import "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/identities/contents"

// Builder represents an instance builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithContent(content contents.Content) Builder
	Now() (Identities, error)
}

// Instance represents an identities command
type Identities interface {
	Name() string
	Content() contents.Content
}
