package contents

import "github.com/steve-care-software/steve/domain/commands/inputs/shares/dashboards/contents/fetches"

// Builder represents the content builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	Now() (Content, error)
}

// Content reresents content
type Content interface {
	IsFetch() bool
	Fetch() fetches.Fetch
}
