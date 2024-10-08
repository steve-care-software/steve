package access

import (
	"errors"

	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/writes"
)

type builder struct {
	write writes.Write
	read  permissions.Permission
}

func createBuilder() Builder {
	out := builder{
		write: nil,
		read:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithWrite adds a write to the builder
func (app *builder) WithWrite(write writes.Write) Builder {
	app.write = write
	return app
}

// WithRead adds a read to the builder
func (app *builder) WithRead(read permissions.Permission) Builder {
	app.read = read
	return app
}

// Now builds a new Access instance
func (app *builder) Now() (Access, error) {
	if app.write == nil {
		return nil, errors.New("the write is mandatory in order to build an Access instance")
	}

	if app.read != nil {
		return createAccessWithRead(app.write, app.read), nil
	}

	return createAccess(app.write), nil
}
