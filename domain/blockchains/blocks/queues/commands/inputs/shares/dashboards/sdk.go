package dashboards

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/shares/dashboards/contents"

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithContent(content contents.Content) Builder
	Now() (Dashboard, error)
}

// Dashboard represents a dashboard
type Dashboard interface {
	Name() string
	Content() contents.Content
}
