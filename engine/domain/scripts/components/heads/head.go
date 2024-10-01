package heads

import (
	"github.com/steve-care-software/steve/commons/hash"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/engine/domain/scripts/components/roles"
)

type head struct {
	hash         hash.Hash
	name         string
	version      uint
	role         roles.Role
	compensation compensations.Compensation
}

func createHead(
	hash hash.Hash,
	name string,
	version uint,
) Head {
	return createHeadInternally(
		hash,
		name,
		version,
		nil,
		nil,
	)
}

func createHeadWithRole(
	hash hash.Hash,
	name string,
	version uint,
	role roles.Role,
) Head {
	return createHeadInternally(
		hash,
		name,
		version,
		role,
		nil,
	)
}

func createHeadWithCompensation(
	hash hash.Hash,
	name string,
	version uint,
	compensation compensations.Compensation,
) Head {
	return createHeadInternally(
		hash,
		name,
		version,
		nil,
		compensation,
	)
}

func createHeadWithRoleAndCompensation(
	hash hash.Hash,
	name string,
	version uint,
	role roles.Role,
	compensation compensations.Compensation,
) Head {
	return createHeadInternally(
		hash,
		name,
		version,
		role,
		compensation,
	)
}

func createHeadInternally(
	hash hash.Hash,
	name string,
	version uint,
	role roles.Role,
	compensation compensations.Compensation,
) Head {
	out := head{
		hash:         hash,
		name:         name,
		version:      version,
		role:         role,
		compensation: compensation,
	}

	return &out
}

// Hash returns the hash
func (obj *head) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *head) Name() string {
	return obj.name
}

// Version returns the version
func (obj *head) Version() uint {
	return obj.version
}

// HasRole returns true if there is a role, false otherwise
func (obj *head) HasRole() bool {
	return obj.role != nil
}

// Role returns the role, if any
func (obj *head) Role() roles.Role {
	return obj.role
}

// HasCompensation returns true if there is a compensation, false otherwise
func (obj *head) HasCompensation() bool {
	return obj.compensation != nil
}

// Compensation returns the compensation, if any
func (obj *head) Compensation() compensations.Compensation {
	return obj.compensation
}
