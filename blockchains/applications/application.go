package applications

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	blockchains "github.com/steve-care-software/steve/blockchains/domain"
	"github.com/steve-care-software/steve/blockchains/domain/blocks"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/blockchains/domain/identities"
	"github.com/steve-care-software/steve/blockchains/domain/roots"
	"github.com/steve-care-software/steve/blockchains/domain/rules"
	"github.com/steve-care-software/steve/engine/applications/cryptography"
	"github.com/steve-care-software/steve/hash"
	lists "github.com/steve-care-software/steve/lists/applications"
	resources "github.com/steve-care-software/steve/resources/applications"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type application struct {
	storeListApp               lists.Application
	resourceApp                resources.Application
	cryptographyApp            cryptography.Application
	identityAdapter            identities.Adapter
	identityBuilder            identities.Builder
	blockchainAdapter          blockchains.Adapter
	blockchainBuilder          blockchains.Builder
	rootBuilder                roots.Builder
	rulesBuilder               rules.Builder
	blocksAdapter              blocks.Adapter
	blocksBuilder              blocks.Builder
	blockBuilder               blocks.BlockBuilder
	contentBuilder             contents.Builder
	transactionsBuilder        transactions.Builder
	transactionBuilder         transactions.TransactionBuilder
	entryBuilder               entries.Builder
	hashAdapter                hash.Adapter
	identityNamesList          string
	blockchainListKeyname      string
	identityKeynamePrefix      string
	identityUnitsKeynamePrefix string
	blockchainKeynamePrefix    string
	scriptKeynamePrefix        string
	blockKeynamePrefix         string
	trxQueue                   []transactions.Transaction
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
		storeListApp:               storeListApp,
		resourceApp:                resourceApp,
		cryptographyApp:            cryptographyApp,
		identityAdapter:            identityAdapter,
		identityBuilder:            identityBuilder,
		blockchainAdapter:          blockchainAdapter,
		blockchainBuilder:          blockchainBuilder,
		rootBuilder:                rootBuilder,
		rulesBuilder:               rulesBuilder,
		blocksAdapter:              blocksAdapter,
		blocksBuilder:              blocksBuilder,
		blockBuilder:               blockBuilder,
		contentBuilder:             contentBuilder,
		transactionsBuilder:        transactionsBuilder,
		transactionBuilder:         transactionBuilder,
		entryBuilder:               entryBuilder,
		hashAdapter:                hashAdapter,
		identityNamesList:          identityNamesList,
		blockchainListKeyname:      blockchainListKeyname,
		identityKeynamePrefix:      identityKeynamePrefix,
		identityUnitsKeynamePrefix: identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:    blockchainKeynamePrefix,
		scriptKeynamePrefix:        scriptKeynamePrefix,
		blockKeynamePrefix:         blockKeynamePrefix,
		trxQueue:                   []transactions.Transaction{},
	}

	return &out
}

// AmountIdentities returns the amount of identities
func (app *application) AmountIdentities() (*uint, error) {
	return app.storeListApp.Amount(app.identityNamesList)
}

// Identities lists the identity names:
func (app *application) Identities(index uint, amount uint) ([]string, error) {
	list, err := app.storeListApp.Retrieve(app.identityNamesList, index, amount)
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
func (app *application) Register(name string, password []byte, language uint8) ([]string, error) {
	cipher, seedWords, err := app.generateIdentityThenEncrypt(name, password, language)
	if err != nil {
		return nil, err
	}

	keyname := fmt.Sprintf("%s%s", app.identityKeynamePrefix, name)
	err = app.resourceApp.Insert(keyname, cipher)
	if err != nil {
		return nil, err
	}

	err = app.storeListApp.Append(app.identityNamesList, [][]byte{
		[]byte(name),
	})

	if err != nil {
		return nil, err
	}

	err = app.resourceApp.Commit()
	if err != nil {
		return nil, err
	}

	return seedWords, nil
}

// Authenticate authenticates in an identity:
func (app *application) Authenticate(name string, password []byte) (identities.Identity, error) {
	keyname := fmt.Sprintf("%s%s", app.identityKeynamePrefix, name)
	cipher, err := app.resourceApp.Retrieve(keyname)
	if err != nil {
		return nil, err
	}

	data, err := app.cryptographyApp.Decrypt(cipher, password)
	if err != nil {
		return nil, err
	}

	identity, _, err := app.identityAdapter.ToInstance(data)
	if err != nil {
		return nil, err
	}

	return identity, nil
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

	return app.resourceApp.Commit()
}

// Units returns the amount of units the authenticated identity has
func (app *application) Units(identity identities.Identity, blockchain uuid.UUID) (*uint64, error) {
	pubKey := identity.PK().Public().(ed25519.PublicKey)
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
func (app *application) Transact(identity identities.Identity, script hash.Hash, fees uint64, flag hash.Hash) error {
	entry, err := app.entryBuilder.Create().
		WithFees(fees).
		WithFlag(flag).
		WithScript(script).
		Now()

	if err != nil {
		return err
	}

	message := entry.Hash().Bytes()
	pk := identity.PK()
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
func (app *application) Mine(identity identities.Identity, blockchainID uuid.UUID, maxAmountTrx uint) error {
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

	minerPubKey := identity.PK().Public()
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
	identity identities.Identity,
	identifier uuid.UUID,
	name string,
	description string,
	unitAmount uint64,
	miningValue uint8,
	baseDifficulty uint8,
	increaseDiffPerrx float64,
) error {
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

	pubKey := identity.PK().Public().(ed25519.PublicKey)
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
	err = app.resourceApp.Insert(unitsKeyname, unitAmountBytes)
	if err != nil {
		return err
	}

	return app.resourceApp.Commit()
}

// AmountBlockchains returns the amount of blockchains
func (app *application) AmountBlockchains() (*uint, error) {
	return app.storeListApp.Amount(app.blockchainListKeyname)
}

// Blockchains returns the list of blockchains
func (app *application) Blockchains(index uint, amount uint) ([]uuid.UUID, error) {
	retBytes, err := app.storeListApp.Retrieve(app.blockchainListKeyname, index, amount)
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
	err = app.resourceApp.Save(blockchainKeyname, blockchainBytes)
	if err != nil {
		return err
	}

	return app.resourceApp.Commit()
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

func (app *application) generateIdentityThenEncrypt(name string, password []byte, language uint8) ([]byte, []string, error) {
	pk, seedWords, err := app.cryptographyApp.GeneratePrivateKey(language)
	if err != nil {
		return nil, nil, err
	}

	return app.generateIdentityFromPKThenEncrypt(name, password, pk, seedWords)
}

func (app *application) generateIdentityFromSeedWordsThenEncrypt(name string, password []byte, seedWords []string) ([]byte, error) {
	pk, err := app.cryptographyApp.GeneratePrivateKeyFromSeedWords(seedWords)
	if err != nil {
		return nil, err
	}

	identity, _, err := app.generateIdentityFromPKThenEncrypt(name, password, pk, seedWords)
	if err != nil {
		return nil, err
	}

	return identity, nil
}

func (app *application) generateIdentityFromPKThenEncrypt(name string, password []byte, pk ed25519.PrivateKey, seedWords []string) ([]byte, []string, error) {
	identity, err := app.identityBuilder.Create().
		WithName(name).
		WithPK(pk).
		Now()

	if err != nil {
		return nil, nil, err
	}

	data, err := app.identityAdapter.ToBytes(identity)
	if err != nil {
		return nil, nil, err
	}

	retBytes, err := app.cryptographyApp.Encrypt(data, password)
	if err != nil {
		return nil, nil, err
	}

	return retBytes, seedWords, nil
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
