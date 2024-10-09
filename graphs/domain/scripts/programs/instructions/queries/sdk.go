package queries

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/routes"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/saves"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/selects"
)

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
