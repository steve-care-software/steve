package queries

import (
	"github.com/google/uuid"
)

type query struct {
	from uuid.UUID
	to   uuid.UUID
}

func createQuery(
	from uuid.UUID,
	to uuid.UUID,
) Query {
	out := query{
		from: from,
		to:   to,
	}

	return &out
}

// From returns the from identifier
func (obj *query) From() uuid.UUID {
	return obj.from
}

// To returns the to identifier
func (obj *query) To() uuid.UUID {
	return obj.to
}
