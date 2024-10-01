package queries

import "github.com/google/uuid"

// NewQueriesForTests creates a new queries for tests
func NewQueriesForTests(list []Query) Queries {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewQueryForTests creates a new query for tests
func NewQueryForTests(from uuid.UUID, to uuid.UUID) Query {
	ins, err := NewQueryBuilder().Create().From(from).To(to).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
