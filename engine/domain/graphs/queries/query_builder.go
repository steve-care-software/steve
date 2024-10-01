package queries

import (
	"errors"

	"github.com/google/uuid"
)

type queryBuilder struct {
	pFrom *uuid.UUID
	pTo   *uuid.UUID
}

func createQueryBuilder() QueryBuilder {
	out := queryBuilder{
		pFrom: nil,
		pTo:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryBuilder) Create() QueryBuilder {
	return createQueryBuilder()
}

// From adds a from identifier to the builder
func (app *queryBuilder) From(from uuid.UUID) QueryBuilder {
	app.pFrom = &from
	return app
}

// To adds a to identifier to the builder
func (app *queryBuilder) To(to uuid.UUID) QueryBuilder {
	app.pTo = &to
	return app
}

// Now builds a new Query instance
func (app *queryBuilder) Now() (Query, error) {
	if app.pFrom == nil {
		return nil, errors.New("the from identifier is mandatory in order to build a Query instance")
	}

	if app.pTo == nil {
		return nil, errors.New("the to identifier is mandatory in order to build a Query instance")
	}

	return createQuery(*app.pFrom, *app.pTo), nil
}
