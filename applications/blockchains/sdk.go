package blockchains

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

const maxDifficulty = 64

// Application represents the blockchain application
type Application interface {
	// Identities lists the identity names:
	Identities() ([]string, error)

	// Register registers a new identity:
	Register(name string, password []byte, seedWords []string) error

	// Authenticate authenticates in an identity:
	Authenticate(name string, password []byte) error

	// Recover recovers an identity using the seed phrases
	Recover(name string, newPassword []byte, words []string) error

	// Authenticated returns the authenticated idgentity, if any
	Authenticated() (string, error)

	// Units returns the amount of units the authenticated identity has
	Units() (*uint64, error)

	// Transact creates a new transaction and adds it to our queue list
	Transact(script []byte, fees uint64, flag hash.Hash) error

	// TrxQueue returns the transactions ready to be put in a block
	TrxQueue() (transactions.Transactions, error)

	// Difficulty speculates the difficulty based on the amount of trx
	Difficulty(blockchainID uuid.UUID, amountTrx uint) (*uint8, error)

	// Mine mines a block using the queued transaction, with the specified max amount of trx
	Mine(blockchain uuid.UUID, maxAmountTrx uint) error

	// BlocksQueue returns the mined blocks queue
	BlocksQueue() (blocks.Blocks, error)

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
