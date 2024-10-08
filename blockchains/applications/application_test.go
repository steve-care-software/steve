package applications

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/engine/applications/cryptography"
	"github.com/steve-care-software/steve/hash"
	lists "github.com/steve-care-software/steve/lists/applications"
	resources "github.com/steve-care-software/steve/resources/applications"
)

func TestApplication_Success(t *testing.T) {
	script := []byte(`
		head:
			engine: v1;
			name: mySchema;
			access: 
				read: .first .second (0.2);
				write: 
					.first .again;
					review: .first .second .third (0.1);
				;
			;
		;

		son;
		father;
		grandFather;
		grandGrandFather;

		father[0,3](son[1,]): .son .father
						| .father .grandFather
						| .grandFather .grandGrandFather
						---
							mySuite[ .mySchema[son] .grandGrandFather]:
								!(.son .father .grandFather .grandGrandFather);
								(.son .father .grandFather .grandGrandFather);
							;
						;

		grandFather(grandSon[2,]): .son .grandFather
								| .father .grandGrandFather
								;
	`)

	baseDir := "./test_files"
	defer func() {
		os.RemoveAll(baseDir)
	}()

	randomID, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	resourceApp, err := resources.NewBuilder().Create().
		WithBasePath(baseDir).
		WithReadChunkSize(1024).
		WithTargetIdentifier("target.tmp").
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	listApp, err := lists.NewBuilder().Create().
		WithResource(resourceApp).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application, err := NewBuilder(
		"identities",
		"blockchains",
		"identities:by_name:",
		"units:by_blockchain_and_pubkeyhash:",
		"blockchain:by_uuid:",
		"block:by_hash:",
	).Create().
		WithResource(resourceApp).
		WithList(listApp).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstUsername := "roger"
	firstPassword := []byte("this is a password")
	seedWords, err := application.Register(firstUsername, firstPassword, cryptography.LangEnglish)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(seedWords) != 24 {
		t.Errorf("%d seed words were expected, %d returned", 24, len(seedWords))
		return
	}

	pIdentitiesAmount, err := application.AmountIdentities()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if *pIdentitiesAmount != 1 {
		t.Errorf("%d identities were expected, %d returned", 1, *pIdentitiesAmount)
		return
	}

	identitiesList, err := application.Identities(0, 1)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(identitiesList) != 1 {
		t.Errorf("%d identities were expected, %d returned", 1, len(identitiesList))
		return
	}

	if identitiesList[0] != firstUsername {
		t.Errorf("the identity at index 0 was expected to be %s, %s returned", firstUsername, identitiesList[0])
		return
	}

	_, err = application.Authenticate(firstUsername, firstPassword)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// bad password
	_, err = application.Authenticate(firstUsername, []byte("bad password"))
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	firstNewPassword := []byte("this is the new password")
	err = application.Recover(firstUsername, firstNewPassword, seedWords)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// old password
	_, err = application.Authenticate(firstUsername, firstPassword)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	retIdentity, err := application.Authenticate(firstUsername, firstNewPassword)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve invalid blockchain units:
	_, err = application.Units(retIdentity, randomID)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	// create a new blockchain:
	blockchainID, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	unitAmount := uint64(100000000000)
	err = application.Create(
		retIdentity,
		blockchainID,
		"myBlockchain",
		"This is a description",
		unitAmount,
		0,
		2,
		0.01,
	)

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pUnitsAmount, err := application.Units(retIdentity, blockchainID)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if unitAmount != *pUnitsAmount {
		t.Errorf("the amount was expected to be %d, %d returned", unitAmount, *pUnitsAmount)
		return
	}

	retBlockchain, err := application.Blockchain(blockchainID)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retBlockchain.Identifier().String() != blockchainID.String() {
		t.Errorf("the blockchain identifier was expetced to be %s, %s returned", retBlockchain.Identifier().String(), blockchainID.String())
		return
	}

	pBlockchainAmount, err := application.AmountBlockchains()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if *pBlockchainAmount != 1 {
		t.Errorf("%d blockchains were expected, %d returned", 1, *pBlockchainAmount)
		return
	}

	retBlockchainIds, err := application.Blockchains(0, 1)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retBlockchainIds) != 1 {
		t.Errorf("%d bockchains were expected, %d returned", 1, len(retBlockchainIds))
		return
	}

	fees := 200
	pFlag, err := hash.NewAdapter().FromBytes([]byte("this is some flag"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Transact(retIdentity, script, uint64(fees), *pFlag)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	trxQueue, err := application.TrxQueue()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	trxQueueList := trxQueue.List()
	if len(trxQueueList) != 1 {
		t.Errorf("the queue was expected to contain %d trx, %d returned", 1, len(trxQueueList))
		return
	}

	// mine:
	err = application.Mine(retIdentity, blockchainID, 2)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve blockchain:
	retBlockchainAfterMine, err := application.Blockchain(blockchainID)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !retBlockchainAfterMine.HasHead() {
		t.Errorf("the blockchain was expected to contain an head")
		return
	}
}
