package blockchains

import (
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
)

// NewBlockchainForTests creates a new blockchain for tests
func NewBlockchainForTests(
	identifier uuid.UUID,
	name string,
	description string,
	rules rules.Rules,
	root roots.Root,
	createdOn time.Time,
) Blockchain {
	ins, err := NewBuilder().Create().
		WithIdentifier(identifier).
		WithName(name).
		WithDescription(description).
		WithRules(rules).
		WithRoot(root).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewBlockchainWithHeadForTests creates a new blockchain with head for tests
func NewBlockchainWithHeadForTests(
	identifier uuid.UUID,
	name string,
	description string,
	rules rules.Rules,
	root roots.Root,
	createdOn time.Time,
	head blocks.Block,
) Blockchain {
	ins, err := NewBuilder().Create().
		WithIdentifier(identifier).
		WithName(name).
		WithDescription(description).
		WithRules(rules).
		WithRoot(root).
		WithHead(head).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
