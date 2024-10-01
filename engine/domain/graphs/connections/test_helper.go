package connections

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections/links"
)

// NewConnectionsForTests creates a new connections for tests
func NewConnectionsForTests(list []Connection) Connections {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConnectionForTests creates a new connection for tests
func NewConnectionForTests(from uuid.UUID, link links.Link, to uuid.UUID) Connection {
	ins, err := NewConnectionBuilder().Create().From(from).To(to).WithLink(link).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
