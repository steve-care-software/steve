package blockchains

import (
	"os"
	"testing"

	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/applications/resources/lists"
)

func TestApplication_Success(t *testing.T) {
	baseDir := "./test_files"
	defer func() {
		os.RemoveAll(baseDir)
	}()

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
		"identities:units:by_name:",
		"blockchain:byt_uuid:",
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
}
