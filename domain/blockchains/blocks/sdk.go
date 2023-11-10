package blocks

import (
	"time"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands"
	"github.com/steve-care-software/steve/domain/blockchains/hash"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Message() string
	Commands() commands.Commands
	Parent() hash.Hash
	CreatedOn() time.Time
}
