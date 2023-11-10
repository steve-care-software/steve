package blocks

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Builder represents the block builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithCommands(commands commands.Commands) Builder
	WithParent(parent hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Message() string
	Commands() commands.Commands
	Parent() hash.Hash
	CreatedOn() time.Time
}
