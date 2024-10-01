package resources

import (
	"bytes"
	"os"
	"testing"
)

func TestApplication_Success(t *testing.T) {
	baseDir := "./test_files"
	defer func() {
		os.RemoveAll(baseDir)
	}()

	application, err := NewBuilder().Create().WithBasePath(baseDir).WithReadChunkSize(1024).WithTargetIdentifier("target.tmp").Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	firstData := []byte("this is some first data")
	thirdData := []byte("this is some third data")
	data := map[string][]byte{
		"first":  firstData,
		"second": []byte("this is some second data"),
		"third":  thirdData,
	}

	for keyname, data := range data {
		err = application.Insert(keyname, data)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}
	}

	err = application.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstData, err := application.Retrieve("first")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(firstData, retFirstData) {
		t.Errorf("the returned data is invalid for identifier: first")
		return
	}

	err = application.Delete("first")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = application.Retrieve("first")
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	retThirdData, err := application.Retrieve("third")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(thirdData, retThirdData) {
		t.Errorf("the returned data is invalid for identifier: third")
		return
	}

	err = application.Rollback(1)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstData, err = application.Retrieve("first")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(firstData, retFirstData) {
		t.Errorf("the returned data is invalid for identifier: first")
		return
	}

	updatedFirst := []byte("this is the first updated data")
	err = application.Save("first", updatedFirst)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retFirstData, err = application.Retrieve("first")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retFirstData, updatedFirst) {
		t.Errorf("the returned data is invalid for identifier: first")
		return
	}

	err = application.Insert("first", []byte("this is some data"))
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
