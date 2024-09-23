package schemas

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/points"
)

type builder struct {
	hashAdapter  hash.Adapter
	name         string
	version      uint
	points       points.Points
	connections  connections.Connections
	role         roles.Role
	compensation compensations.Compensation
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		name:         "",
		version:      0,
		points:       nil,
		connections:  nil,
		role:         nil,
		compensation: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.version = version
	return app
}

// WithPoints adds a points to the builder
func (app *builder) WithPoints(points points.Points) Builder {
	app.points = points
	return app
}

// WithConnections adds a connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// WithRole adds a role to the builder
func (app *builder) WithRole(role roles.Role) Builder {
	app.role = role
	return app
}

// WithCompensation adds a compensation to the builder
func (app *builder) WithCompensation(compensation compensations.Compensation) Builder {
	app.compensation = compensation
	return app
}

// Now builds a new Schema instance
func (app *builder) Now() (Schema, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Schema instance")
	}

	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build a Schema instance")
	}

	if app.points == nil {
		return nil, errors.New("the points is mandatory in order to build a Schema instance")
	}

	if app.connections == nil {
		return nil, errors.New("the connections is mandatory in order to build a Schema instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(app.version))),
		app.points.Hash().Bytes(),
		app.connections.Hash().Bytes(),
	}

	if app.role != nil {
		data = append(data, app.role.Hash().Bytes())
	}

	if app.compensation != nil {
		data = append(data, app.compensation.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.role != nil && app.compensation != nil {
		return createSchemaWithRoleAndCompensation(
			*pHash,
			app.name,
			app.version,
			app.points,
			app.connections,
			app.role,
			app.compensation,
		), nil
	}

	if app.role != nil {
		return createSchemaWithRole(
			*pHash,
			app.name,
			app.version,
			app.points,
			app.connections,
			app.role,
		), nil
	}

	if app.compensation != nil {
		return createSchemaWithCompensation(
			*pHash,
			app.name,
			app.version,
			app.points,
			app.connections,
			app.compensation,
		), nil
	}

	return createSchema(
		*pHash,
		app.name,
		app.version,
		app.points,
		app.connections,
	), nil
}
