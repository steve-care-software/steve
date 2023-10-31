package updates

import "github.com/steve-care-software/steve/domain/commands/administrators/inputs/administrators/instances/contents/updates/contents"

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithObject(object string) Builder
	WithContent(content contents.Content) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Object() string
	Content() contents.Content
}
