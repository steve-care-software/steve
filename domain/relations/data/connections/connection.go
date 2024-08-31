package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/relations/data/connections/links"
)

type connection struct {
	from uuid.UUID
	link links.Link
	to   uuid.UUID
}

func createConnection(
	from uuid.UUID,
	link links.Link,
	to uuid.UUID,
) Connection {
	out := connection{
		from: from,
		link: link,
		to:   to,
	}

	return &out
}

// From returns the from identifier
func (obj *connection) From() uuid.UUID {
	return obj.from
}

// Link returns the link
func (obj *connection) Link() links.Link {
	return obj.link
}

// To returns the to identifier
func (obj *connection) To() uuid.UUID {
	return obj.to
}
