package blockchains

import (
	"crypto"
	"crypto/ed25519"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/applications/cryptography"
	"github.com/steve-care-software/steve/applications/stores/lists"
	"github.com/steve-care-software/steve/applications/stores/resources"
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

type application struct {
	cryptographyApp              cryptography.Application
	storeListApp                 lists.Application
	resourceApp                  resources.Application
	identityAdapter              identities.Adapter
	identityBuilder              identities.Builder
	blockchainBuilder            blockchains.Builder
	rootBuilder                  roots.Builder
	rulesBuilder                 rules.Builder
	blockBuilder                 blocks.Builder
	contentBuilder               contents.Builder
	transactionsBuilder          transactions.Builder
	transactionBuilder           transactions.TransactionBuilder
	entryBuilder                 entries.Builder
	identityNamesList            string
	identityKeynamePrefix        string
	identityUnitsKeynamePrefix   string
	blockchainKeynamePrefix      string
	currentAuthenticatedIdentity identities.Identity
	trxQueue                     []transactions.Transaction
}

func createApplication(
	cryptographyApp cryptography.Application,
	storeListApp lists.Application,
	resourceApp resources.Application,
	identityAdapter identities.Adapter,
	identityBuilder identities.Builder,
	blockchainBuilder blockchains.Builder,
	rootBuilder roots.Builder,
	rulesBuilder rules.Builder,
	blockBuilder blocks.Builder,
	contentBuilder contents.Builder,
	transactionsBuilder transactions.Builder,
	transactionBuilder transactions.TransactionBuilder,
	entryBuilder entries.Builder,
	identityNamesList string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
) Application {
	out := application{
		cryptographyApp:              cryptographyApp,
		storeListApp:                 storeListApp,
		resourceApp:                  resourceApp,
		identityAdapter:              identityAdapter,
		identityBuilder:              identityBuilder,
		blockchainBuilder:            blockchainBuilder,
		rootBuilder:                  rootBuilder,
		rulesBuilder:                 rulesBuilder,
		blockBuilder:                 blockBuilder,
		contentBuilder:               contentBuilder,
		transactionsBuilder:          transactionsBuilder,
		transactionBuilder:           transactionBuilder,
		entryBuilder:                 entryBuilder,
		identityNamesList:            identityNamesList,
		identityKeynamePrefix:        identityKeynamePrefix,
		identityUnitsKeynamePrefix:   identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:      blockchainKeynamePrefix,
		currentAuthenticatedIdentity: nil,
		trxQueue:                     []transactions.Transaction{},
	}

	return &out
}

// Identities lists the identity names:
func (app *application) Identities() ([]string, error) {
	list, err := app.storeListApp.RetrieveAll(app.identityNamesList)
	if err != nil {
		return nil, err
	}

	output := []string{}
	for _, oneName := range list {
		output = append(output, string(oneName))
	}

	return output, nil
}

// Register registers a new identity:
func (app *application) Register(name string, password []byte, seedWords []string) error {
	cipher, err := app.generateIdentityFromSeedWordsThenEncrypt(name, password, seedWords)
	if err != nil {
		return err
	}

	keyname := fmt.Sprintf("%s%s", app.identityKeynamePrefix, name)
	return app.resourceApp.Insert(keyname, cipher)
}

// Authenticate authenticates in an identity:
func (app *application) Authenticate(name string, password []byte) error {
	keyname := fmt.Sprintf("%s%s", app.identityKeynamePrefix, name)
	cipher, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return err
	}

	data, err := app.cryptographyApp.Decrypt(cipher, password)
	if err != nil {
		return err
	}

	identity, err := app.identityAdapter.ToInstance(data)
	if err != nil {
		return err
	}

	app.currentAuthenticatedIdentity = identity
	return nil
}

// Recover recovers an identity using the seed phrases
func (app *application) Recover(name string, newPassword []byte, words []string) error {
	cipher, err := app.generateIdentityFromSeedWordsThenEncrypt(name, newPassword, words)
	if err != nil {
		return err
	}

	keyname := fmt.Sprintf("%s%s", app.identityKeynamePrefix, name)
	return app.resourceApp.Insert(keyname, cipher)
}

// Authenticated returns the authenticated idgentity, if any
func (app *application) Authenticated() (string, error) {
	return app.currentAuthenticatedIdentity.Name(), nil
}

// Units returns the amount of units the authenticated identity hasg
func (app *application) Units() (*uint64, error) {
	if app.currentAuthenticatedIdentity != nil {
		return nil, errors.New("there is no authenticated identity")
	}

	keyname := fmt.Sprintf("%s%s", app.identityUnitsKeynamePrefix, app.currentAuthenticatedIdentity.Name())
	data, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	amount := binary.LittleEndian.Uint64(data)
	return &amount, nil
}

// Transact creates a new transaction and adds it to our queue list
func (app *application) Transact(blockchain uuid.UUID, script []byte, fees uint64, flag hash.Hash) error {
	if app.currentAuthenticatedIdentity != nil {
		return errors.New("there is no authenticated identity")
	}

	keyname := fmt.Sprintf("%s%s", app.blockchainKeynamePrefix, blockchain.String())
	_, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return err
	}

	entry, err := app.entryBuilder.Create().
		WithFees(fees).
		WithFlag(flag).
		WithScript(script).
		Now()

	if err != nil {
		return err
	}

	message := entry.Hash().Bytes()
	signature, err := app.currentAuthenticatedIdentity.PK().Sign(nil, message, crypto.SHA512)
	if err != nil {
		return err
	}

	trx, err := app.transactionBuilder.Create().WithEntry(entry).WithSignature(signature).Now()
	if err != nil {
		return err
	}

	app.trxQueue = append(app.trxQueue, trx)
	return nil
}

// Queue returns the transactions ready to be put in a block
func (app *application) Queue() (transactions.Transactions, error) {
	if len(app.trxQueue) <= 0 {
		return nil, errors.New("there is currently no transaction in queue")
	}

	return app.transactionsBuilder.Create().
		WithList(app.trxQueue).
		Now()
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

func (app *application) generateIdentityFromSeedWordsThenEncrypt(name string, password []byte, seedWords []string) ([]byte, error) {
	seed := []byte{}
	for _, oneWord := range seedWords {
		seed = append(seed, []byte(oneWord)...)
	}

	pk := ed25519.NewKeyFromSeed(seed)
	identity, err := app.identityBuilder.Create().
		WithName(name).
		WithPK(pk).
		Now()

	if err != nil {
		return nil, err
	}

	data, err := app.identityAdapter.ToBytes(identity)
	if err != nil {
		return nil, err
	}

	return app.cryptographyApp.Encrypt(data, password)
}
