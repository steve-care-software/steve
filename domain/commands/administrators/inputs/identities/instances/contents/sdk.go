package contents

import (
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/identities/instances/contents/deletes"
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/identities/instances/contents/fetches"
	"github.com/steve-care-software/steve/domain/commands/administrators/inputs/identities/instances/contents/inserts"
)

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(del deletes.Delete) Builder
	Now() (Content, error)
}

// Content reresents content
type Content interface {
	IsFetch() bool
	Fetch() fetches.Fetch
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
}
