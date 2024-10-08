package heads

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/engine/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	name         string
	version      uint
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

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Head instance")
	}

	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build a Head instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(app.version))),
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
		return createHeadWithRoleAndCompensation(
			*pHash,
			app.name,
			app.version,
			app.role,
			app.compensation,
		), nil
	}

	if app.role != nil {
		return createHeadWithRole(
			*pHash,
			app.name,
			app.version,
			app.role,
		), nil
	}

	if app.compensation != nil {
		return createHeadWithCompensation(
			*pHash,
			app.name,
			app.version,
			app.compensation,
		), nil
	}

	return createHead(
		*pHash,
		app.name,
		app.version,
	), nil
}
