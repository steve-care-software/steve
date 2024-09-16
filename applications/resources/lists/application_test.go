package lists

import (
	"os"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/applications/resources"
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

	application, err := NewBuilder().Create().
		WithResource(resourceApp).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	identiier := "myData"
	values := [][]byte{
		[]byte("this is first element"),
		[]byte("this is second element"),
		[]byte("this is third element"),
	}

	err = application.Append(identiier, values)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pAmount, err := application.Amount(identiier)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if *pAmount != uint(len(values)) {
		t.Errorf("%d values were expected, %d returned", *pAmount, len(values))
		return
	}

	retValues, err := application.Retrieve(identiier, 1, 1)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(values[1:2], retValues) {
		t.Errorf("the returned values are invalid")
		return
	}

	_, err = application.Retrieve(identiier, uint(len(values)), 1)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	_, err = application.Retrieve(identiier, 0, uint(len(values)))
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

	otherData := [][]byte{
		[]byte("this is some addit data"),
		[]byte("yesm, agaib"),
	}

	err = application.Append(identiier, otherData)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Commit()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	expected := append(values, otherData...)
	retValues, err = application.RetrieveAll(identiier)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(expected, retValues) {
		t.Errorf("the returned values are invalid")
		return
	}

}
