package successes

import "github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/shares/dashboards/successes/fetches"

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	IsFetch() bool
	Fetch() fetches.Fetch
}
