package connections

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/graphs/connections/links"
	"github.com/steve-care-software/steve/domain/hash"
)

type connection struct {
	hash hash.Hash
	from uuid.UUID
	link links.Link
	to   uuid.UUID
}

func createConnection(
	hash hash.Hash,
	from uuid.UUID,
	link links.Link,
	to uuid.UUID,
) Connection {
	out := connection{
		hash: hash,
		from: from,
		link: link,
		to:   to,
	}

	return &out
}

// Hash returns the hash
func (obj *connection) Hash() hash.Hash {
	return obj.hash
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

// Debug returns the string debug representation of the link
func (obj *connection) Debug() string {
	return fmt.Sprintf("\nfrom: %s\nto: %s\nlink: %s\n-----\n", obj.from.String(), obj.to.String(), obj.link.Debug())
}
