package blockchains

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
	"github.com/steve-care-software/steve/domain/hash"
)

type blockchain struct {
	identifier  uuid.UUID
	name        string
	description string
	rules       rules.Rules
	root        roots.Root
	head        blocks.Block
	createdOn   time.Time
}

func createBlockchain(
	identifier uuid.UUID,
	name string,
	description string,
	rules rules.Rules,
	root roots.Root,
	createdOn time.Time,
) Blockchain {
	return createBlockchainInternally(
		identifier,
		name,
		description,
		rules,
		root,
		createdOn,
		nil,
	)
}

func createBlockchainWithHead(
	identifier uuid.UUID,
	name string,
	description string,
	rules rules.Rules,
	root roots.Root,
	createdOn time.Time,
	head blocks.Block,
) Blockchain {
	return createBlockchainInternally(
		identifier,
		name,
		description,
		rules,
		root,
		createdOn,
		head,
	)
}

func createBlockchainInternally(
	identifier uuid.UUID,
	name string,
	description string,
	rules rules.Rules,
	root roots.Root,
	createdOn time.Time,
	head blocks.Block,
) Blockchain {
	out := blockchain{
		identifier:  identifier,
		name:        name,
		description: description,
		rules:       rules,
		root:        root,
		createdOn:   createdOn,
		head:        head,
	}

	return &out
}

// Identifier returns the identifier
func (obj *blockchain) Identifier() uuid.UUID {
	return obj.identifier
}

// Name returns the name
func (obj *blockchain) Name() string {
	return obj.name
}

// Description returns the description
func (obj *blockchain) Description() string {
	return obj.description
}

// Rules returns the rules
func (obj *blockchain) Rules() rules.Rules {
	return obj.rules
}

// Root returns the rules
func (obj *blockchain) Root() roots.Root {
	return obj.root
}

// CreatedOn returns the creation time
func (obj *blockchain) CreatedOn() time.Time {
	return obj.createdOn
}

// Difficulty returns the difficulty needed for the provided amount of trx
func (obj *blockchain) Difficulty(amountTrx uint) (*uint8, error) {
	baseDifficulty := uint64(obj.rules.BaseDifficulty())
	increateDiffPerTrx := obj.rules.IncreaseDifficultyPerTrx()
	incrAmount := uint64(increateDiffPerTrx * float64(amountTrx))
	difficulty := baseDifficulty + incrAmount
	if difficulty > hash.Size {
		str := fmt.Sprintf("the max difficulty amount was expected to at max %d, %d calculated", hash.Size, difficulty)
		return nil, errors.New(str)
	}

	casted := uint8(difficulty)
	return &casted, nil
}

// HasHead returns true if there is an head, false otherwise
func (obj *blockchain) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *blockchain) Head() blocks.Block {
	return obj.head
}
