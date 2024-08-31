package connections

import (
	"errors"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/relations/data/connections/links"
)

type connectionBuilder struct {
	pFrom *uuid.UUID
	link  links.Link
	pTo   *uuid.UUID
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		pFrom: nil,
		link:  nil,
		pTo:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
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
		return nil, errors.New("the from is mandatory in order to build a Connection instance")
	}

	if app.pTo == nil {
		return nil, errors.New("the to is mandatory in order to build a Connection instance")
	}

	return createConnection(*app.pFrom, app.link, *app.pTo), nil
}
