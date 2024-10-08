package entries

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/hash"
)

func TestAdapter_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pFlag, err := hashAdapter.FromBytes([]byte("flag"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pScript, err := hashAdapter.FromBytes([]byte("script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	entry := NewEntryForTests(*pFlag, *pScript, 34)
	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(entry)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retEntry, retRemaining, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(entry, retEntry) {
		t.Errorf("the returned entry is invalid")
		return
	}
}

func TestAdapter_withRemaining_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pFlag, err := hashAdapter.FromBytes([]byte("flag"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pScript, err := hashAdapter.FromBytes([]byte("script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	entry := NewEntryForTests(*pFlag, *pScript, 34)
	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(entry)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retEntry, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(entry, retEntry) {
		t.Errorf("the returned entry is invalid")
		return
	}
}
