package blockchains

import (
	"crypto/ed25519"
	"encoding/base64"
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
	identityNamesList            string
	blockchainListKeyname        string
	identityKeynamePrefix        string
	identityUnitsKeynamePrefix   string
	blockchainKeynamePrefix      string
	scriptKeynamePrefix          string
	blockKeynamePrefix           string
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
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
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
		identityNamesList:            identityNamesList,
		blockchainListKeyname:        blockchainListKeyname,
		identityKeynamePrefix:        identityKeynamePrefix,
		identityUnitsKeynamePrefix:   identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:      blockchainKeynamePrefix,
		scriptKeynamePrefix:          scriptKeynamePrefix,
		blockKeynamePrefix:           blockKeynamePrefix,
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
	keyname := app.unitsPerOwnerAndBlockchainKeyname(pubKey, blockchain)
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
func (app *application) Transact(script hash.Hash, fees uint64, flag hash.Hash) error {
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
	pk := app.currentAuthenticatedIdentity.PK()
	signature := ed25519.Sign(pk, message)
	publicKey := pk.Public().(ed25519.PublicKey)
	trx, err := app.transactionBuilder.Create().WithEntry(entry).
		WithSignature(signature).
		WithPublicKey(publicKey).
		Now()

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

// Mine mines a block using the queued transaction, with the specified max amount of trx
func (app *application) Mine(blockchainID uuid.UUID, maxAmountTrx uint) error {
	if app.currentAuthenticatedIdentity == nil {
		return errors.New(noAuthIdentityErr)
	}

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

	pDifficulty, err := blockchain.Difficulty(uint(len(trxList)))
	if err != nil {
		return err
	}

	trx, err := app.transactionsBuilder.Create().WithList(trxList).Now()
	if err != nil {
		return err
	}

	rules := blockchain.Rules()
	miningValue := rules.MiningValue()
	result, err := mine(app.hashAdapter, trx, *pDifficulty, miningValue)
	if err != nil {
		return err
	}

	parent := blockchain.Root().Hash()
	if blockchain.HasHead() {
		parent = blockchain.Head().Hash()
	}

	head, err := app.resourceApp.Head()
	if err != nil {
		return err
	}

	minerPubKey := app.currentAuthenticatedIdentity.PK().Public()
	content, err := app.contentBuilder.Create().
		WithParent(parent).
		WithTransactions(trx).
		WithMiner(minerPubKey.(ed25519.PublicKey)).
		WithCommit(head.Hash()).
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

	app.trxQueue = remaining
	return app.block(blockchain, block)
}

// Block adds a block as the head of its blockchain
func (app *application) Block(blockchainID uuid.UUID, block blocks.Block) error {
	blockchain, err := app.Blockchain(blockchainID)
	if err != nil {
		return err
	}

	return app.block(blockchain, block)
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

	head, err := app.resourceApp.Head()
	if err != nil {
		return err
	}

	pubKey := app.currentAuthenticatedIdentity.PK().Public().(ed25519.PublicKey)
	root, err := app.rootBuilder.Create().
		WithAmount(unitAmount).
		WithOwner(pubKey).
		WithCommit(head.Hash()).
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
	unitsKeyname := app.unitsPerOwnerAndBlockchainKeyname(pubKey, identifier)
	return app.resourceApp.Insert(unitsKeyname, unitAmountBytes)
}

// Blockchains returns the list of blockchains
func (app *application) Blockchains() ([]uuid.UUID, error) {
	retBytes, err := app.storeListApp.RetrieveAll(app.blockchainListKeyname)
	if err != nil {
		return nil, err
	}

	list := []uuid.UUID{}
	for _, oneUUIDBytes := range retBytes {
		uuid, err := uuid.FromBytes(oneUUIDBytes)
		if err != nil {
			return nil, err
		}

		list = append(list, uuid)
	}

	return list, nil
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

func (app *application) block(blockchain blockchains.Blockchain, block blocks.Block) error {
	err := app.validateBlock(blockchain, block)
	if err != nil {
		return err
	}

	pCommit, err := app.resourceApp.Head()
	if err != nil {
		return err
	}

	commitHash := pCommit.Hash()
	blockCommitHash := block.Content().Commit()
	if !blockCommitHash.Compare(commitHash) {
		err := app.resourceApp.RollbackTo(blockCommitHash)
		if err != nil {
			return err
		}
	}

	retBytes, err := app.blocksAdapter.InstanceToBytes(block)
	if err != nil {
		return err
	}

	err = app.transferFees(blockchain, block)
	if err != nil {
		return err
	}

	keyname := fmt.Sprintf("%s%s", app.blockKeynamePrefix, block.Hash().String())
	err = app.resourceApp.Insert(keyname, retBytes)
	if err != nil {
		return err
	}

	identifier := blockchain.Identifier()
	name := blockchain.Name()
	description := blockchain.Description()
	rules := blockchain.Rules()
	root := blockchain.Root()
	createdOn := blockchain.CreatedOn()
	updated, err := app.blockchainBuilder.Create().
		WithIdentifier(identifier).
		WithName(name).
		WithDescription(description).
		WithRules(rules).
		WithRoot(root).
		WithHead(block).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		return err
	}

	blockchainBytes, err := app.blockchainAdapter.ToBytes(updated)
	if err != nil {
		return err
	}

	blockchainKeyname := fmt.Sprintf("%s%s", app.blockchainKeynamePrefix, blockchain.Identifier().String())
	return app.resourceApp.Save(blockchainKeyname, blockchainBytes)
}

func (app *application) validateBlock(blockchain blockchains.Blockchain, block blocks.Block) error {
	if !blockchain.HasHead() {
		return nil
	}

	parent := blockchain.Head().Hash()
	miningValue := blockchain.Rules().MiningValue()
	pCurrentWork, err := app.canculateWorkFromHeadToBlock(miningValue, parent, block)
	if err != nil {
		return err
	}

	pBlockWork, err := app.workFromBlock(miningValue, block)
	if err != nil {
		return err
	}

	// replace the block:
	if *pCurrentWork < *pBlockWork {
		return nil
	}

	str := fmt.Sprintf(
		"the provided block (hash: %s) contains %d difficulty of work, its blockchain (id: %s) contains %d difficulty of work after that block, therefore the provided block cannot be added to the blockchain",
		block.Hash().String(),
		*pBlockWork,
		blockchain.Identifier().String(),
		*pCurrentWork,
	)

	return errors.New(str)
}

func (app *application) canculateWorkFromHeadToBlock(miningValue uint8, headHash hash.Hash, block blocks.Block) (*uint, error) {
	keyname := fmt.Sprintf("%s%s", app.blockKeynamePrefix, headHash.String())
	blockBytes, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	head, _, err := app.blocksAdapter.BytesToInstance(blockBytes)
	if err != nil {
		return nil, err
	}

	if head.Hash().Compare(block.Hash()) {
		output := uint(0)
		return &output, nil
	}

	pHeadWork, err := app.workFromBlock(miningValue, head)
	if err != nil {
		return nil, err
	}

	parentBlockHash := head.Content().Parent()
	pWork, err := app.canculateWorkFromHeadToBlock(miningValue, parentBlockHash, block)
	if err != nil {
		return nil, err
	}

	output := *pHeadWork + *pWork
	return &output, nil
}

func (app *application) workFromBlock(miningValue uint8, block blocks.Block) (*uint, error) {
	result := block.Result()
	trxHash := block.Content().Transactions().Hash()
	pHash, err := executeHash(app.hashAdapter, trxHash, result)
	if err != nil {
		return nil, err
	}

	difficulty := uint(0)
	hashBytes := pHash.Bytes()
	for _, oneByte := range hashBytes {
		if oneByte != miningValue {
			break
		}

		difficulty++
	}

	return &difficulty, nil
}

func (app *application) transferFees(blockchain blockchains.Blockchain, block blocks.Block) error {
	blockchainID := blockchain.Identifier()
	content := block.Content()
	minerRevenue := uint64(0)
	transactionsList := content.Transactions().List()
	for _, oneTrx := range transactionsList {
		pubKey := oneTrx.PublicKey()
		keyname := app.unitsPerOwnerAndBlockchainKeyname(pubKey, blockchainID)

		walletAmount := uint64(0)
		unitAmountBytes, err := app.resourceApp.Retrieve(keyname)
		if err != nil {
			return err
		}

		pWalletAmount, err := pointers.BytesToUint64(unitAmountBytes)
		if err == nil {
			walletAmount = (*pWalletAmount)
		}

		trxFees := oneTrx.Entry().Fees()
		if walletAmount < trxFees {
			str := fmt.Sprintf("the transaction (hash: %s) of block (hash: %s) in blockchain (id: %s) was expected to pay %d units in fees, but the wallet only contains %d units", oneTrx.Hash().String(), block.Hash().String(), blockchainID.String(), trxFees, walletAmount)
			return errors.New(str)
		}

		updatedAmount := walletAmount - trxFees
		updatedAmountBytes := pointers.Uint64ToBytes(updatedAmount)
		err = app.resourceApp.Save(keyname, updatedAmountBytes)
		if err != nil {
			return err
		}

		minerRevenue += trxFees
	}

	// retrieve the miner wallet amount:
	miner := content.Miner()
	keyname := app.unitsPerOwnerAndBlockchainKeyname(miner, blockchainID)
	unitAmountBytes, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return err
	}

	currentAmount := uint64(0)
	pWalletAmount, err := pointers.BytesToUint64(unitAmountBytes)
	if err == nil {
		currentAmount = *pWalletAmount
	}

	updatedAmount := currentAmount + minerRevenue
	updatedAmountBytes := pointers.Uint64ToBytes(updatedAmount)
	return app.resourceApp.Save(keyname, updatedAmountBytes)
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

func (app *application) unitsPerOwnerAndBlockchainKeyname(owner ed25519.PublicKey, blockchain uuid.UUID) string {
	encoded := base64.StdEncoding.EncodeToString(owner)
	return fmt.Sprintf("%s%s%s", app.identityUnitsKeynamePrefix, blockchain.String(), encoded)
}
