package actions

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/actions/deletes"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"
)

// Builder represents an actions builder
type Builder interface {
	Create() Builder
	WithList(list []Action) Builder
	Now() (Actions, error)
}

// Actions represents actions
type Actions interface {
	List() []Action
}

// ActionBuilder represents an action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithInsert(insert resources.Resource) ActionBuilder
	WithDelete(del deletes.Delete) ActionBuilder
	Now() (Action, error)
}

// Action represents a commit action
type Action interface {
	IsInsert() bool
	Insert() resources.Resource
	IsDelete() bool
	Delete() deletes.Delete
}
