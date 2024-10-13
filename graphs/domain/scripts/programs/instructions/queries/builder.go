package queries

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/bridges"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/routes"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/saves"
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/selects"
)

type builder struct {
	save   saves.Save
	sel    selects.Select
	bridge bridges.Bridge
	route  routes.Route
}

func createBuilder() Builder {
	return &builder{
		save:   nil,
		sel:    nil,
		bridge: nil,
		route:  nil,
	}
}

// Create initializes the query builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSave adds a save to the query builder
func (app *builder) WithSave(save saves.Save) Builder {
	app.save = save
	return app
}

// WithSelect adds a select to the query builder
func (app *builder) WithSelect(sel selects.Select) Builder {
	app.sel = sel
	return app
}

// WithBridge adds a bridge to the query builder
func (app *builder) WithBridge(bridge bridges.Bridge) Builder {
	app.bridge = bridge
	return app
}

// WithRoute adds a route to the query builder
func (app *builder) WithRoute(route routes.Route) Builder {
	app.route = route
	return app
}

// Now builds a new Query instance
func (app *builder) Now() (Query, error) {
	if app.save != nil {
		return createQueryWithSave(app.save), nil
	}

	if app.sel != nil {
		return createQueryWithSelect(app.sel), nil
	}

	if app.bridge != nil {
		return createQueryWithBridge(app.bridge), nil
	}

	if app.route != nil {
		return createQueryWithRoute(app.route), nil
	}

	return nil, errors.New("the Query is invalid")
}
