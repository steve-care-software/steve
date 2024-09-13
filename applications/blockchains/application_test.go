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

	err = application.Register("roger", []byte("myPassword"), seedWords)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
