package administrators

import "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/shares/administrators/creates"

// Administrator represents an administrator
type Administrator interface {
	IsCreate() bool
	Create() creates.Create
}
