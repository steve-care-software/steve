package blockchains

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/applications/cryptography"
	resources "github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/applications/resources/lists"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/blockchains/identities"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
	"github.com/steve-care-software/steve/domain/hash"
)

const noAuthIdentityErr = "there is currently no authenticated identity"

// NewBuilder creates a new builder
func NewBuilder(
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
	blockQueueKeyname string,
) Builder {
	cryptographyApp := cryptography.NewApplication()
	identityAdapter := identities.NewAdapter()
	identityBuilder := identities.NewBuilder()
	blockchainAdapter := blockchains.NewAdapter()
	blockchainBuilder := blockchains.NewBuilder()
	rootBuilder := roots.NewBuilder()
	rulesBuilder := rules.NewBuilder()
	blocksAdapter := blocks.NewAdapter()
	blocksBuilder := blocks.NewBuilder()
	blockBuilder := blocks.NewBlockBuilder()
	contentBuilder := contents.NewBuilder()
	transactionsBuilder := transactions.NewBuilder()
	transactionBuilder := transactions.NewTransactionBuilder()
	entryBuilder := entries.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		cryptographyApp,
		identityAdapter,
		identityBuilder,
		blockchainAdapter,
		blockchainBuilder,
		rootBuilder,
		rulesBuilder,
		blocksAdapter,
		blocksBuilder,
		blockBuilder,
		contentBuilder,
		transactionsBuilder,
		transactionBuilder,
		entryBuilder,
		hashAdapter,
		identityNamesList,
		blockchainListKeyname,
		identityKeynamePrefix,
		identityUnitsKeynamePrefix,
		blockchainKeynamePrefix,
		scriptKeynamePrefix,
		blockKeynamePrefix,
		blockQueueKeyname,
	)
}

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithResource(resourceApp resources.Application) Builder
	WithList(listApp lists.Application) Builder
	Now() (Application, error)
}

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
	Units(blockchain uuid.UUID) (*uint64, error)

	// Transact creates a new transaction and adds it to our queue list
	Transact(script hash.Hash, fees uint64, flag hash.Hash) error

	// TrxQueue returns the transactions ready to be put in a block
	TrxQueue() (transactions.Transactions, error)

	// Mine mines a block using the queued transaction, with the specified max amount of trx
	Mine(blockchain uuid.UUID, maxAmountTrx uint) error

	// Block adds a block to the queue
	Block(blockchain uuid.UUID, block blocks.Block) error

	// BlocksQueue returns the mined blocks queue
	BlocksQueue() (blocks.Blocks, error)

	// Sync syncs the mined blocks with the network
	Sync(blockHash hash.Hash) error

	// Create a new blockchain
	Create(identifier uuid.UUID, name string, description string, unitAmount uint64, miningValue uint8, baseDifficulty uint8, increaseDiffPerrx float64) error

	// Blockchains returns the list of blockchains
	Blockchains() ([]uuid.UUID, error)

	// Blockchain returns the blochain by id
	Blockchain(identifier uuid.UUID) (blockchains.Blockchain, error)

	// Script returns the script by its hash
	Script(hash hash.Hash) ([]byte, error)
}
