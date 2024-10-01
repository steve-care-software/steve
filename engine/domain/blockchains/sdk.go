package blockchains

import (
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/engine/domain/blockchains/roots"
	"github.com/steve-care-software/steve/engine/domain/blockchains/rules"
)

const dataLengthTooSmallErrPattern = "(blockchain) the data length was expected to be at least %d bytes, %d returned"
const uuidSize = 16

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	blockAdapter := blocks.NewAdapter()
	rulesAdapter := rules.NewAdapter()
	rootAdapter := roots.NewAdapter()
	builder := NewBuilder()
	return createAdapter(
		blockAdapter,
		rulesAdapter,
		rootAdapter,
		builder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the blockchain adapter
type Adapter interface {
	ToBytes(ins Blockchain) ([]byte, error)
	ToInstance(data []byte) (Blockchain, []byte, error)
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
	Difficulty(amountTrx uint) (*uint8, error)
	HasHead() bool
	Head() blocks.Block
}
