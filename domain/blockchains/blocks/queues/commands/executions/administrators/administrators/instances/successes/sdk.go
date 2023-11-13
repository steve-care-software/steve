package successes

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/fetches"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/administrators/administrators/instances/successes/updates"
)

// Builder represents the success builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithUpdate(update updates.Update) Builder
	WithDelete(delete deletes.Delete) Builder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	IsFetch() bool
	Fetch() fetches.Fetch
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() deletes.Delete
}
