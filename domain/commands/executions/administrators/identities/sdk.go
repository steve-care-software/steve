package identities

import (
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/deletes"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/fetches"
	"github.com/steve-care-software/steve/domain/commands/executions/administrators/identities/inserts"
)

// Builder represents an identities builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(del deletes.Delete) Builder
	Now() (Identities, error)
}

// Identities represents identities
type Identities interface {
	IsFetch() bool
	Fetch() fetches.Fetch
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
}
