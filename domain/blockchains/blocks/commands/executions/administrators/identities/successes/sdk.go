package successes

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/fetches"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/executions/administrators/identities/successes/inserts"
)

// Builder represents a success builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(del deletes.Delete) Builder
	Now() (Success, error)
}

// Success represents success
type Success interface {
	IsFetch() bool
	Fetch() fetches.Fetch
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
}
