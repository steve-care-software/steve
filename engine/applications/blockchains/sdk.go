package blockchains

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/applications/cryptography"
	"github.com/steve-care-software/steve/engine/domain/blockchains"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/engine/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/engine/domain/blockchains/identities"
	"github.com/steve-care-software/steve/engine/domain/blockchains/roots"
	"github.com/steve-care-software/steve/engine/domain/blockchains/rules"
	"github.com/steve-care-software/steve/commons/hash"
	resources resources "github.com/steve-care-software/steve/databases/resources/applications"
	lists "github.com/steve-care-software/steve/databases/lists/applications"
)

// NewBuilder creates a new builder
func NewBuilder(
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
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
	// AmountIdentities returns the amount of identities
	AmountIdentities() (*uint, error)

	// Identities lists the identity names:
	Identities(index uint, amount uint) ([]string, error)

	// Register registers a new identity:
	Register(name string, password []byte, language uint8) ([]string, error)

	// Authenticate authenticates in an identity:
	Authenticate(name string, password []byte) (identities.Identity, error)

	// Recover recovers an identity using the seed phrases
	Recover(name string, newPassword []byte, words []string) error

	// Units returns the amount of units the authenticated identity has
	Units(identity identities.Identity, blockchain uuid.UUID) (*uint64, error)

	// Transact creates a new transaction and adds it to our queue list
	Transact(identity identities.Identity, script hash.Hash, fees uint64, flag hash.Hash) error

	// TrxQueue returns the transactions ready to be put in a block
	TrxQueue() (transactions.Transactions, error)

	// Mine mines a block using the queued transaction, with the specified max amount of trx
	Mine(identity identities.Identity, blockchain uuid.UUID, maxAmountTrx uint) error

	// Block adds a block to the queue
	Block(blockchain uuid.UUID, block blocks.Block) error

	// Sync syncs the mined blocks with the network
	Sync(blockHash hash.Hash) error

	// Create a new blockchain
	Create(identity identities.Identity, identifier uuid.UUID, name string, description string, unitAmount uint64, miningValue uint8, baseDifficulty uint8, increaseDiffPerrx float64) error

	// AmountBlockchains returns the amount of blockchains
	AmountBlockchains() (*uint, error)

	// Blockchains returns the list of blockchains
	Blockchains(index uint, amount uint) ([]uuid.UUID, error)

	// Blockchain returns the blochain by id
	Blockchain(identifier uuid.UUID) (blockchains.Blockchain, error)

	// Script returns the script by its hash
	Script(hash hash.Hash) ([]byte, error)
}
