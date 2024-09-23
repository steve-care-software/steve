package contexts

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/compensations"
	"github.com/steve-care-software/steve/domain/scripts/components/roles"
	"github.com/steve-care-software/steve/domain/scripts/specifics/contexts/contents"
)

type context struct {
	hash         hash.Hash
	name         string
	version      uint
	content      contents.Content
	parent       string
	role         roles.Role
	compensation compensations.Compensation
}

func createContextWithParentAndRoleAndCompensation(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	parent string,
	role roles.Role,
	compensation compensations.Compensation,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		parent,
		role,
		compensation,
	)
}

func createContextWithParentAndCompensation(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	parent string,
	compensation compensations.Compensation,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		parent,
		nil,
		compensation,
	)
}

func createContextWithParentAndRole(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	parent string,
	role roles.Role,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		parent,
		role,
		nil,
	)
}

func createContextWithRoleAndCompensation(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	role roles.Role,
	compensation compensations.Compensation,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		"",
		role,
		compensation,
	)
}

func createContextWithParent(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	parent string,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		parent,
		nil,
		nil,
	)
}

func createContextWithRole(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	role roles.Role,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		"",
		role,
		nil,
	)
}

func createContextWithCompensation(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	compensation compensations.Compensation,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		"",
		nil,
		compensation,
	)
}

func createContext(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
) Context {
	return createContextInternally(
		hash,
		name,
		version,
		content,
		"",
		nil,
		nil,
	)
}

func createContextInternally(
	hash hash.Hash,
	name string,
	version uint,
	content contents.Content,
	parent string,
	role roles.Role,
	compensation compensations.Compensation,
) Context {
	out := context{
		hash:         hash,
		name:         name,
		version:      version,
		content:      content,
		parent:       parent,
		role:         role,
		compensation: compensation,
	}

	return &out
}

// Hash returns the hash
func (obj *context) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *context) Name() string {
	return obj.name
}

// Version returns the version
func (obj *context) Version() uint {
	return obj.version
}

// Content returns the content
func (obj *context) Content() contents.Content {
	return obj.content
}

// HasParent returns true if there is a parent, false otherwise
func (obj *context) HasParent() bool {
	return obj.parent != ""
}

// Parent returns the parent, if any
func (obj *context) Parent() string {
	return obj.parent
}

// HasRole returns true if there is a role, false otherwise
func (obj *context) HasRole() bool {
	return obj.role != nil
}

// Role returns the role, if any
func (obj *context) Role() roles.Role {
	return obj.role
}

// HasCompensation returns true if there is a compensation, false otherwise
func (obj *context) HasCompensation() bool {
	return obj.compensation != nil
}

// Compensation returns the compensation, if any
func (obj *context) Compensation() compensations.Compensation {
	return obj.compensation
}
