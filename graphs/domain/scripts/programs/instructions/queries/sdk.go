package queries

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/routes"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/saves"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/selects"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithSave(save saves.Save) Builder
	WithSelect(sel selects.Select) Builder
	WithBridge(bridge bridges.Bridge) Builder
	WithRoute(route routes.Route) Builder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	IsSave() bool
	Save() saves.Save
	IsSelect() bool
	Select() selects.Select
	IsBridge() bool
	Bridge() bridges.Bridge
	IsRoute() bool
	Route() routes.Route
}
