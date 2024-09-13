package blockchains

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/applications/resources/lists"
)

func TestApplication_Success(t *testing.T) {
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
		"script:by_hash:",
		"block:by_hash:",
		"block_queues",
	).Create().WithResource(resourceApp).WithList(listApp).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	seedWords := []string{
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"abandon",
		"about",
	}

	firstUsername := "roger"
	firstPassword := []byte("this is a password")
	err = application.Register(firstUsername, firstPassword, seedWords)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	identitiesList, err := application.Identities()
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

	_, err = application.Authenticated()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = application.Authenticate(firstUsername, firstPassword)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// bad password
	err = application.Authenticate(firstUsername, []byte("bad password"))
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	authUsername, err := application.Authenticated()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if authUsername != firstUsername {
		t.Errorf("the authenticated username was expected to be %s, %s returned", authUsername, firstPassword)
		return
	}

	firstNewPassword := []byte("this is the new password")
	err = application.Recover(firstUsername, firstNewPassword, seedWords)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// old password
	err = application.Authenticate(firstUsername, firstPassword)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = application.Authenticate(firstUsername, firstNewPassword)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	// retrieve invalid blockchain units:
	_, err = application.Units(randomID)
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

	err = resourceApp.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pAmount, err := application.Units(blockchainID)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if unitAmount != *pAmount {
		t.Errorf("the amount was expected to be %d, %d returned", unitAmount, *pAmount)
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

	retBlockchainIds, err := application.Blockchains()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retBlockchainIds) != 1 {
		t.Errorf("%d bockchains were expected, %d returned", len(retBlockchainIds), 1)
		return
	}

}
