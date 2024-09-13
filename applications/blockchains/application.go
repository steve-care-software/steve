package blockchains

import (
	"crypto"
	"crypto/ed25519"
	"errors"
	"fmt"
	"time"

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
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
	"github.com/steve-care-software/steve/domain/uuids"
)

type application struct {
	storeListApp                 lists.Application
	resourceApp                  resources.Application
	cryptographyApp              cryptography.Application
	identityAdapter              identities.Adapter
	identityBuilder              identities.Builder
	blockchainAdapter            blockchains.Adapter
	blockchainBuilder            blockchains.Builder
	rootBuilder                  roots.Builder
	rulesBuilder                 rules.Builder
	blocksAdapter                blocks.Adapter
	blocksBuilder                blocks.Builder
	blockBuilder                 blocks.BlockBuilder
	contentBuilder               contents.Builder
	transactionsBuilder          transactions.Builder
	transactionBuilder           transactions.TransactionBuilder
	entryBuilder                 entries.Builder
	hashAdapter                  hash.Adapter
	uuidAdapter                  uuids.Adapter
	identityNamesList            string
	blockchainListKeyname        string
	identityKeynamePrefix        string
	identityUnitsKeynamePrefix   string
	blockchainKeynamePrefix      string
	scriptKeynamePrefix          string
	blockKeynamePrefix           string
	blockQueueKeyname            string
	currentAuthenticatedIdentity identities.Identity
	trxQueue                     []transactions.Transaction
}

func createApplication(
	storeListApp lists.Application,
	resourceApp resources.Application,
	cryptographyApp cryptography.Application,
	identityAdapter identities.Adapter,
	identityBuilder identities.Builder,
	blockchainAdapter blockchains.Adapter,
	blockchainBuilder blockchains.Builder,
	rootBuilder roots.Builder,
	rulesBuilder rules.Builder,
	blocksAdapter blocks.Adapter,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	contentBuilder contents.Builder,
	transactionsBuilder transactions.Builder,
	transactionBuilder transactions.TransactionBuilder,
	entryBuilder entries.Builder,
	hashAdapter hash.Adapter,
	uuidAdapter uuids.Adapter,
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
	blockQueueKeyname string,
) Application {
	out := application{
		storeListApp:                 storeListApp,
		resourceApp:                  resourceApp,
		cryptographyApp:              cryptographyApp,
		identityAdapter:              identityAdapter,
		identityBuilder:              identityBuilder,
		blockchainAdapter:            blockchainAdapter,
		blockchainBuilder:            blockchainBuilder,
		rootBuilder:                  rootBuilder,
		rulesBuilder:                 rulesBuilder,
		blocksAdapter:                blocksAdapter,
		blocksBuilder:                blocksBuilder,
		blockBuilder:                 blockBuilder,
		contentBuilder:               contentBuilder,
		transactionsBuilder:          transactionsBuilder,
		transactionBuilder:           transactionBuilder,
		entryBuilder:                 entryBuilder,
		hashAdapter:                  hashAdapter,
		uuidAdapter:                  uuidAdapter,
		identityNamesList:            identityNamesList,
		blockchainListKeyname:        blockchainListKeyname,
		identityKeynamePrefix:        identityKeynamePrefix,
		identityUnitsKeynamePrefix:   identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:      blockchainKeynamePrefix,
		scriptKeynamePrefix:          scriptKeynamePrefix,
		blockKeynamePrefix:           blockKeynamePrefix,
		blockQueueKeyname:            blockQueueKeyname,
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
	err = app.resourceApp.Insert(keyname, cipher)
	if err != nil {
		return err
	}

	return app.storeListApp.Append(app.identityNamesList, [][]byte{
		[]byte(name),
	})
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

	identity, _, err := app.identityAdapter.ToInstance(data)
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
	err = app.resourceApp.Save(keyname, cipher)
	if err != nil {
		return err
	}

	app.currentAuthenticatedIdentity = nil
	return nil
}

// Authenticated returns the authenticated idgentity, if any
func (app *application) Authenticated() (string, error) {
	if app.currentAuthenticatedIdentity == nil {
		return "", errors.New(noAuthIdentityErr)
	}

	return app.currentAuthenticatedIdentity.Name(), nil
}

// Units returns the amount of units the authenticated identity has
func (app *application) Units(blockchain uuid.UUID) (*uint64, error) {
	if app.currentAuthenticatedIdentity == nil {
		return nil, errors.New(noAuthIdentityErr)
	}

	pubKey := app.currentAuthenticatedIdentity.PK().Public().(ed25519.PublicKey)
	pubKeyHash, err := app.hashAdapter.FromBytes(pubKey)
	if err != nil {
		return nil, err
	}

	keyname := app.unitsPerOwnerAndBlockchainKeyname(*pubKeyHash, blockchain)
	data, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	pAmount, err := pointers.BytesToUint64(data)
	if err != nil {
		return nil, err
	}

	return pAmount, nil
}

// Transact creates a new transaction and adds it to our queue list
func (app *application) Transact(script []byte, fees uint64, flag hash.Hash) error {
	if app.currentAuthenticatedIdentity == nil {
		return errors.New(noAuthIdentityErr)
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

// TrxQueue returns the transactions ready to be put in a block
func (app *application) TrxQueue() (transactions.Transactions, error) {
	if len(app.trxQueue) <= 0 {
		return nil, errors.New("there is currently no transaction in queue")
	}

	return app.transactionsBuilder.Create().
		WithList(app.trxQueue).
		Now()
}

// Difficulty speculates the difficulty based on the amount of trx
func (app *application) Difficulty(blockchainID uuid.UUID, amountTrx uint) (*uint8, error) {
	blockchain, err := app.retrieveBlockchainFromID(blockchainID)
	if err != nil {
		return nil, err
	}

	return app.difficulty(blockchain, amountTrx)
}

// Mine mines a block using the queued transaction, with the specified max amount of trx
func (app *application) Mine(blockchainID uuid.UUID, maxAmountTrx uint) error {
	blockchain, err := app.retrieveBlockchainFromID(blockchainID)
	if err != nil {
		return err
	}

	trxList := app.trxQueue
	remaining := []transactions.Transaction{}
	length := uint(len(app.trxQueue))
	if length > maxAmountTrx {
		trxList = app.trxQueue[0:maxAmountTrx]
		remaining = app.trxQueue[maxAmountTrx:]
	}

	pDifficulty, err := app.difficulty(blockchain, uint(len(trxList)))
	if err != nil {
		return err
	}

	result, err := mine(trxList, *pDifficulty)
	if err != nil {
		return err
	}

	parent := blockchain.Root().Hash()
	if blockchain.HasHead() {
		parent = blockchain.Head().Hash()
	}

	transactions, err := app.transactionsBuilder.Create().WithList(trxList).Now()
	if err != nil {
		return err
	}

	content, err := app.contentBuilder.Create().
		WithParent(parent).
		WithTransactions(transactions).
		Now()

	if err != nil {
		return err
	}

	block, err := app.blockBuilder.Create().
		WithContent(content).
		WithResult(result).
		Now()

	if err != nil {
		return err
	}

	retBytes, err := app.blocksAdapter.InstanceToBytes(block)
	if err != nil {
		return err
	}

	keyname := fmt.Sprintf("%s%s", app.blockKeynamePrefix, block.Hash().String())
	err = app.resourceApp.Insert(keyname, retBytes)
	if err != nil {
		return err
	}

	retQueueBytes, err := app.resourceApp.Retrieve(app.blockQueueKeyname)
	if err != nil {
		return err
	}

	retQueue, _, err := app.blocksAdapter.BytesToInstances(retQueueBytes)
	if err != nil {
		return err
	}

	queueList := retQueue.List()
	queueList = append(queueList, block)
	blocks, err := app.blocksBuilder.Create().WithList(queueList).Now()
	if err != nil {
		return err
	}

	retBlocksBytes, err := app.blocksAdapter.InstancesToBytes(blocks)
	if err != nil {
		return err
	}

	app.trxQueue = remaining
	return app.resourceApp.Insert(app.blockQueueKeyname, retBlocksBytes)
}

// BlocksQueue returns the mined blocks
func (app *application) BlocksQueue() (blocks.Blocks, error) {
	retBytes, err := app.resourceApp.Retrieve(app.blockQueueKeyname)
	if err != nil {
		return nil, err
	}

	ins, _, err := app.blocksAdapter.BytesToInstances(retBytes)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// Sync syncs the mined blocks with the network
func (app *application) Sync(blockHash hash.Hash) error {
	return nil
}

// Create a new blockchain
func (app *application) Create(
	identifier uuid.UUID,
	name string,
	description string,
	unitAmount uint64,
	miningValue uint8,
	baseDifficulty uint8,
	increaseDiffPerrx float64,
) error {
	if app.currentAuthenticatedIdentity == nil {
		return errors.New(noAuthIdentityErr)
	}

	rules, err := app.rulesBuilder.Create().
		WithBaseDifficulty(baseDifficulty).
		WithIncreaseDifficultyPerTrx(increaseDiffPerrx).
		WithMiningValue(miningValue).
		Now()

	if err != nil {
		return err
	}

	pubKey := app.currentAuthenticatedIdentity.PK().Public().(ed25519.PublicKey)
	pPubKeyHash, err := app.hashAdapter.FromBytes(pubKey)
	if err != nil {
		return err
	}

	root, err := app.rootBuilder.Create().
		WithAmount(unitAmount).
		WithOwner(*pPubKeyHash).
		Now()

	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	blockchain, err := app.blockchainBuilder.Create().
		WithIdentifier(identifier).
		WithName(name).
		WithDescription(description).
		WithRules(rules).
		WithRoot(root).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	retBytes, err := app.blockchainAdapter.ToBytes(blockchain)
	if err != nil {
		return err
	}

	keyname := fmt.Sprintf("%s%s", app.blockchainKeynamePrefix, blockchain.Identifier().String())
	err = app.resourceApp.Insert(keyname, retBytes)
	if err != nil {
		return err
	}

	identifierBytes, err := identifier.MarshalBinary()
	if err != nil {
		return err
	}

	err = app.storeListApp.Append(app.blockchainListKeyname, [][]byte{
		identifierBytes,
	})

	if err != nil {
		return err
	}

	unitAmountBytes := pointers.Uint64ToBytes(unitAmount)
	unitsKeyname := app.unitsPerOwnerAndBlockchainKeyname(*pPubKeyHash, identifier)
	return app.resourceApp.Insert(unitsKeyname, unitAmountBytes)
}

// Blockchains returns the list of blockchains
func (app *application) Blockchains() ([]uuid.UUID, error) {
	retBytes, err := app.resourceApp.Retrieve(app.blockchainListKeyname)
	if err != nil {
		return nil, err
	}

	return app.uuidAdapter.FromBytes(retBytes)
}

// Blockchain returns the blochain by id
func (app *application) Blockchain(identifier uuid.UUID) (blockchains.Blockchain, error) {
	keyname := fmt.Sprintf("%s%s", app.blockchainKeynamePrefix, identifier.String())
	retBytes, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	ins, _, err := app.blockchainAdapter.ToInstance(retBytes)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

// Script returns the script by its hash
func (app *application) Script(hash hash.Hash) ([]byte, error) {
	keyname := fmt.Sprintf("%s%s", app.scriptKeynamePrefix, hash.String())
	return app.resourceApp.Retrieve(keyname)
}

func (app *application) generateIdentityFromSeedWordsThenEncrypt(name string, password []byte, seedWords []string) ([]byte, error) {
	pk, err := app.cryptographyApp.GeneratePrivateKey(seedWords)
	if err != nil {
		return nil, err
	}

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

func (app *application) retrieveBlockchainFromID(id uuid.UUID) (blockchains.Blockchain, error) {
	keyname := fmt.Sprintf("%s%s", app.blockchainKeynamePrefix, id.String())
	retData, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	ins, _, err := app.blockchainAdapter.ToInstance(retData)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (app *application) difficulty(blockchain blockchains.Blockchain, amountTrx uint) (*uint8, error) {
	rules := blockchain.Rules()
	baseDifficulty := uint64(rules.BaseDifficulty())
	increateDiffPerTrx := rules.IncreaseDifficultyPerTrx()
	incrAmount := uint64(increateDiffPerTrx * float64(amountTrx))
	difficulty := baseDifficulty + incrAmount
	if difficulty > maxDifficulty {
		str := fmt.Sprintf("the max difficulty amount was expected to at max %d, %d calculated", maxDifficulty, difficulty)
		return nil, errors.New(str)
	}

	casted := uint8(difficulty)
	return &casted, nil
}

func (app *application) unitsPerOwnerAndBlockchainKeyname(owner hash.Hash, blockchain uuid.UUID) string {
	return fmt.Sprintf("%s%s%s", app.identityUnitsKeynamePrefix, blockchain.String(), owner.String())
}
