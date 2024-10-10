package queries

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/routes"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/saves"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/selects"
)

type query struct {
	save   saves.Save
	sel    selects.Select
	bridge bridges.Bridge
	route  routes.Route
}

func createQueryWithSave(
	save saves.Save,
) Query {
	return createQueryInternally(save, nil, nil, nil)
}

func createQueryWithSelect(
	sel selects.Select,
) Query {
	return createQueryInternally(nil, sel, nil, nil)
}

func createQueryWithBridge(
	bridge bridges.Bridge,
) Query {
	return createQueryInternally(nil, nil, bridge, nil)
}

func createQueryWithRoute(
	route routes.Route,
) Query {
	return createQueryInternally(nil, nil, nil, route)
}

func createQueryInternally(
	save saves.Save,
	sel selects.Select,
	bridge bridges.Bridge,
	route routes.Route,
) Query {
	return &query{
		save:   save,
		sel:    sel,
		bridge: bridge,
		route:  route,
	}
}

// IsSave returns true if the query contains a save
func (obj *query) IsSave() bool {
	return obj.save != nil
}

// Save returns the save of the query
func (obj *query) Save() saves.Save {
	return obj.save
}

// IsSelect returns true if the query contains a select
func (obj *query) IsSelect() bool {
	return obj.sel != nil
}

// Select returns the select of the query
func (obj *query) Select() selects.Select {
	return obj.sel
}

// IsBridge returns true if the query contains a bridge
func (obj *query) IsBridge() bool {
	return obj.bridge != nil
}

// Bridge returns the bridge of the query
func (obj *query) Bridge() bridges.Bridge {
	return obj.bridge
}

// IsRoute returns true if the query contains a route
func (obj *query) IsRoute() bool {
	return obj.route != nil
}

// Route returns the route of the query
func (obj *query) Route() routes.Route {
	return obj.route
}
