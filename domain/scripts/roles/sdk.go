package roles

import "github.com/steve-care-software/steve/domain/hash"

// Role represents the role
type Role interface {
	Version() uint
	Name() string
	HasInsert() bool
	Insert() []hash.Hash
	HasDelete() bool
	Delete() []hash.Hash
}
