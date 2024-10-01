package roles

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	version     uint
	name        string
	insert      []hash.Hash
	del         []hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		version:     0,
		name:        "",
		insert:      nil,
		del:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.version = version
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithInsert adds an insert to the builder
func (app *builder) WithInsert(insert []hash.Hash) Builder {
	app.insert = insert
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(del []hash.Hash) Builder {
	app.del = del
	return app
}

// Now builds a new Role instance
func (app *builder) Now() (Role, error) {
	if app.version <= 0 {
		return nil, errors.New("the version must be greater than zero (0) in order to build a Transfer instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Role instance")
	}

	if app.insert != nil && len(app.insert) <= 0 {
		app.insert = nil
	}

	if app.del != nil && len(app.del) <= 0 {
		app.del = nil
	}

	data := [][]byte{
		[]byte(strconv.Itoa(int(app.version))),
		[]byte(app.name),
	}

	if app.insert != nil {
		for _, oneInsert := range app.insert {
			data = append(data, oneInsert.Bytes())
		}
	}

	if app.del != nil {
		for _, oneDel := range app.del {
			data = append(data, oneDel.Bytes())
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil && app.del != nil {
		return createRoleWithInsertAndDelete(
			*pHash,
			app.version,
			app.name,
			app.insert,
			app.del,
		), nil
	}

	if app.insert != nil {
		return createRoleWithInsert(
			*pHash,
			app.version,
			app.name,
			app.insert,
		), nil
	}

	if app.del != nil {
		return createRoleWithDelete(
			*pHash,
			app.version,
			app.name,
			app.del,
		), nil
	}

	return nil, errors.New("the Role is invalid")
}
