package databases

import (
	"os"
	"testing"
)

func TestExists_thenCreate_thenDelete_Success(t *testing.T) {
	dirPath := "./test_files"
	dstExtension := "destination"
	bckExtension := "backup"
	defer func() {
		os.RemoveAll(dirPath)
	}()

	database := NewApplication(dirPath, dstExtension, bckExtension, nil)

	name := "my_name"
	exists, err := database.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the database was expected to NOT exists")
		return
	}

	err = database.New(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err = database.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !exists {
		t.Errorf("the database was expected to exists")
		return
	}

	err = database.New(name)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	err = database.Delete(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	exists, err = database.Exists(name)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if exists {
		t.Errorf("the database was expected to NOT exists")
		return
	}
}
