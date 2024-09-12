package uuids

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestAdapter_Success(t *testing.T) {
	first, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	second, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	third, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	list := []uuid.UUID{
		first,
		second,
		third,
	}

	adapter := NewAdapter()
	retBytes, err := adapter.FromInstances(list)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retBytes = append(retBytes, []byte("this")...) // some bytes that should be rejected
	retList, err := adapter.FromBytes(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(list, retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}
