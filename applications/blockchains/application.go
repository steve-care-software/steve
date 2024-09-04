package blockchains

import (
	"github.com/google/uuid"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Identities lists the identity names:
func (app *application) Identities() ([]string, error) {
	return nil, nil
}

// Register registers a new identity:
func (app *application) Register(name string, password []byte) ([]string, error) {
	return nil, nil
}

// Authenticate authenticates in an identity:
func (app *application) Authenticate(name string, password []byte) error {
	return nil
}

// Recover recovers an identity using the seed phrases
func (app *application) Recover(name string, words []string, newPassword []byte) error {
	return nil
}

// Authenticated returns the authenticated idgentity, if any
func (app *application) Authenticated() (string, error) {
	return "", nil
}

// Units returns the amount of units the authenticated identity has
func (app *application) Units() (*uint64, error) {
	return nil, nil
}

// Transact creates a new transaction and adds it to our queue list
func (app *application) Transact(blockchain uuid.UUID, script []byte, fees uint64) error {
	return nil
}

// Queue returns the transactions ready to be put in a block
func (app *application) Queue() (transactions.Transactions, error) {
	return nil, nil
}

// Difficulty speculates the difficulty based on the amount of trx
func (app *application) Difficulty(amountTrx uint) (*uint, error) {
	return nil, nil
}

// Mine mines a block using the queued transaction, with the specified max amount of trx
func (app *application) Mine(maxAmountTrx uint) (blocks.Block, error) {
	return nil, nil
}

// Blocks returns the mined blocks
func (app *application) Blocks() ([]hash.Hash, error) {
	return nil, nil
}

// Sync syncs the mined blocks with the network
func (app *application) Sync(blockHash hash.Hash) error {
	return nil
}

// Create a new blockchain
func (app *application) Create(name string, description string, unitAmount uint64, miningValue uint8, baseDifficulty uint8, increaseDiffPerrx float64) error {
	return nil
}

// Blockchains returns the list of blockchains
func (app *application) Blockchains() ([]uuid.UUID, error) {
	return nil, nil
}

// Blockchain returns the blochain by id
func (app *application) Blockchain(identifier uuid.UUID) (blockchains.Blockchain, error) {
	return nil, nil
}

// Script returns the script by its hash
func (app *application) Script(hash hash.Hash) ([]byte, error) {
	return nil, nil
}
