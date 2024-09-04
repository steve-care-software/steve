package blockchains

import (
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier uuid.UUID) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithRules(rules rules.Rules) Builder
	WithRoot(root roots.Root) Builder
	WithHead(head blocks.Block) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Blockchain, error)
}

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
