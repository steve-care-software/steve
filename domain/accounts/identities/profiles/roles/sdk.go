package roles

import "github.com/steve-care-software/steve/domain/accounts/identities/profiles/roles/permissions"

// Role represents a role
type Role interface {
	Name() string
	HasPermissions() bool
	Permissions() permissions.Permissions
}
