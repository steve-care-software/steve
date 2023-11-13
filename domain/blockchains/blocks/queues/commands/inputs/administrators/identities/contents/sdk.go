package contents

import (
	"github.com/steve-care-software/steve/domain/accounts/administrators/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities/contents/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities/contents/fetches"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/administrators/identities/contents/inserts"
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
	Insert() identities.Identity
	IsDelete() bool
	Delete() deletes.Delete
}
