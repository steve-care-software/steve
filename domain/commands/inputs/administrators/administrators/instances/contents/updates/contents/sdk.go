package contents

import "github.com/steve-care-software/steve/domain/commands/inputs/administrators/administrators/instances/contents/updates/contents/passwords"

// Builder represents the content builder
type Builder interface {
	Create() Builder
	WithPassword(password passwords.Password) Builder
	Now() (Content, error)
}

// Content represents the update content
type Content interface {
	IsPassword() bool
	Password() passwords.Password
}
