package blockchains

import (
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
)

// Blockchain represents a blockchain
type Blockchain interface {
	Identifier() uuid.UUID
	Name() string
	Description() string
	Rules() rules.Rules
	Root() roots.Root
	CreatedOn() time.Time
	HasHead() bool
	Head() blocks.Block
}
