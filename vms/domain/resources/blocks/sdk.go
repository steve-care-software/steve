package blocks

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/vms/children/commands/domain/resources"
)

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithCommands(commands resources.Resources) Builder
	WithParent(parent hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Message() string
	Commands() resources.Resources
	Parent() hash.Hash
	CreatedOn() time.Time
}
