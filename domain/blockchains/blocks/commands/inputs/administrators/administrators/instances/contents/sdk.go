package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances/contents/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances/contents/fetches"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/administrators/administrators/instances/contents/updates"
)

// Builder represents a content builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithUpdate(update updates.Update) Builder
	WithDelete(del deletes.Delete) Builder
	Now() (Content, error)
}

// Content reresents content
type Content interface {
	IsFetch() bool
	Fetch() fetches.Fetch
	IsUpdate() bool
	Update() updates.Update
	IsDelete() bool
	Delete() deletes.Delete
}
