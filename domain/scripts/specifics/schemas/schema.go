package schemas

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/connections"
	"github.com/steve-care-software/steve/domain/scripts/specifics/schemas/points"
)

type schema struct {
	hash         hash.Hash
	name         string
	version      uint
	points       points.Points
	connections  connections.Connections
	role         roles.Role
	compensation compensations.Compensation
}

func createSchema(
	hash hash.Hash,
	name string,
	version uint,
	points points.Points,
	connections connections.Connections,
) Schema {
	return createSchemaInternally(
		hash,
		name,
		version,
		points,
		connections,
		nil,
		nil,
	)
}

func createSchemaWithRole(
	hash hash.Hash,
	name string,
	version uint,
	points points.Points,
	connections connections.Connections,
	role roles.Role,
) Schema {
	return createSchemaInternally(
		hash,
		name,
		version,
		points,
		connections,
		role,
		nil,
	)
}

func createSchemaWithCompensation(
	hash hash.Hash,
	name string,
	version uint,
	points points.Points,
	connections connections.Connections,
	compensation compensations.Compensation,
) Schema {
	return createSchemaInternally(
		hash,
		name,
		version,
		points,
		connections,
		nil,
		compensation,
	)
}

func createSchemaWithRoleAndCompensation(
	hash hash.Hash,
	name string,
	version uint,
	points points.Points,
	connections connections.Connections,
	role roles.Role,
	compensation compensations.Compensation,
) Schema {
	return createSchemaInternally(
		hash,
		name,
		version,
		points,
		connections,
		role,
		compensation,
	)
}

func createSchemaInternally(
	hash hash.Hash,
	name string,
	version uint,
	points points.Points,
	connections connections.Connections,
	role roles.Role,
	compensation compensations.Compensation,
) Schema {
	out := schema{
		hash:         hash,
		name:         name,
		version:      version,
		points:       points,
		connections:  connections,
		role:         role,
		compensation: compensation,
	}

	return &out
}

// Hash returns the hash
func (obj *schema) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *schema) Name() string {
	return obj.name
}

// Version returns the version
func (obj *schema) Version() uint {
	return obj.version
}

// Points returns the points
func (obj *schema) Points() points.Points {
	return obj.points
}

// Connections returns the connections
func (obj *schema) Connections() connections.Connections {
	return obj.connections
}

// HasRole returns true if there is a role, false otherwise
func (obj *schema) HasRole() bool {
	return obj.role != nil
}

// Role returns the role, if any
func (obj *schema) Role() roles.Role {
	return obj.role
}

// HasCompensation returns true if there is a compensation, false otherwise
func (obj *schema) HasCompensation() bool {
	return obj.compensation != nil
}

// Compensation returns the compensation, if any
func (obj *schema) Compensation() compensations.Compensation {
	return obj.compensation
}
