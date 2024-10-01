package connections

import (
	"errors"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections/links"
	"github.com/steve-care-software/steve/engine/domain/hash"
)

type connectionBuilder struct {
	hashAdapter hash.Adapter
	pFrom       *uuid.UUID
	link        links.Link
	pTo         *uuid.UUID
}

func createConnectionBuilder(
	hashAdapter hash.Adapter,
) ConnectionBuilder {
	out := connectionBuilder{
		hashAdapter: hashAdapter,
		pFrom:       nil,
		link:        nil,
		pTo:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder(
		app.hashAdapter,
	)
}

// WithLink adds a link to the builder
func (app *connectionBuilder) WithLink(link links.Link) ConnectionBuilder {
	app.link = link
	return app
}

// From adds a from to the builder
func (app *connectionBuilder) From(from uuid.UUID) ConnectionBuilder {
	app.pFrom = &from
	return app
}

// To adds a to to the builder
func (app *connectionBuilder) To(to uuid.UUID) ConnectionBuilder {
	app.pTo = &to
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Connection instance")
	}

	if app.pFrom == nil {
		return nil, errors.New("the from identifier is mandatory in order to build a Connection instance")
	}

	if app.pTo == nil {
		return nil, errors.New("the to identifier is mandatory in order to build a Connection instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		[]byte(app.pFrom.String()),
		[]byte(app.pTo.String()),
	})

	if err != nil {
		return nil, err
	}

	return createConnection(
		*pHash,
		*app.pFrom,
		app.link,
		*app.pTo,
	), nil
}
