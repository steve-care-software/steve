package blockchains

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

// Application represents the blockchain application
type Application interface {
	// Identities lists the identity names:
	Identities() ([]string, error)

	// Register registers a new identity:
	Register(name string, password []byte) ([]string, error)

	// Authenticate authenticates in an identity:
	Authenticate(name string, password []byte) error

	// Recover recovers an identity using the seed phrases
	Recover(name string, words []string, newPassword []byte) error

	// Authenticated returns the authenticated identity, if any
	Authenticated() (string, error)

	// Transact creates a new transaction and adds it to our queue list
	Transact(blockchain uuid.UUID, script []byte, fees uint64) error

	// Queue returns the transactions ready to be put in a block
	Queue() (transactions.Transactions, error)

	// Difficulty speculates the difficulty based on the amount of trx
	Difficulty(amountTrx uint) (*uint, error)

	// Mine mines a block using the queued transaction, with the specified max amount of trx
	Mine(maxAmountTrx uint) (blocks.Block, error)

	// Blocks returns the mined blocks
	Blocks() ([]hash.Hash, error)

	// Sync syncs the mined blocks with the network
	Sync(blockHash hash.Hash) error

	// Create a new blockchain
	Create(name string, description string, unitAmount uint64, miningValue uint8, baseDifficulty uint8, increaseDiffPerrx float64) error

	// Blockchains returns the list of blockchains
	Blockchains() ([]uuid.UUID, error)

	// Blockchain returns the blochain by id
	Blockchain(identifier uuid.UUID) (blockchains.Blockchain, error)

	// Script returns the script by its hash
	Script(hash hash.Hash) ([]byte, error)
}
