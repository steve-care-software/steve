package contexts

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/compensations"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents"
	"github.com/steve-care-software/steve/domain/scripts/contexts/roles"
)

type builder struct {
	hashAdapter  hash.Adapter
	name         string
	version      uint
	content      contents.Content
	parent       string
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
		content:      nil,
		parent:       "",
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

// WithContent adds a content to the builder
func (app *builder) WithContent(content contents.Content) Builder {
	app.content = content
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent string) Builder {
	app.parent = parent
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

// Now builds a new Context instance
func (app *builder) Now() (Context, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Context instance")
	}

	if app.version <= 0 {
		return nil, errors.New("the version is mandatory in order to build a Context instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Context instance")
	}

	data := [][]byte{
		[]byte(app.name),
		[]byte(strconv.Itoa(int(app.version))),
		app.compensation.Hash().Bytes(),
	}

	if app.parent != "" {
		data = append(data, []byte(app.parent))
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

	if app.parent != "" && app.role != nil && app.compensation != nil {
		return createContextWithParentAndRoleAndCompensation(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.parent,
			app.role,
			app.compensation,
		), nil
	}

	if app.parent != "" && app.role != nil {
		return createContextWithParentAndRole(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.parent,
			app.role,
		), nil
	}

	if app.parent != "" && app.compensation != nil {
		return createContextWithParentAndCompensation(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.parent,
			app.compensation,
		), nil
	}

	if app.role != nil && app.compensation != nil {
		return createContextWithRoleAndCompensation(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.role,
			app.compensation,
		), nil
	}

	if app.parent != "" {
		return createContextWithParent(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.parent,
		), nil
	}

	if app.role != nil {
		return createContextWithRole(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.role,
		), nil
	}

	if app.compensation != nil {
		return createContextWithCompensation(
			*pHash,
			app.name,
			app.version,
			app.content,
			app.compensation,
		), nil
	}

	return createContext(*pHash, app.name, app.version, app.content), nil
}
