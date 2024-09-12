package lists

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	list := [][]byte{
		[]byte("this is first"),
		[]byte("this is second data"),
		[]byte("this is another last third data"),
	}

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(list)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retList, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}

}
