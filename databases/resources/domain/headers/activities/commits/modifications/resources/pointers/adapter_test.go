package pointers

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	pointer := NewPointerForTests(0, 23)
	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPointer, retRemaining, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(pointer, retPointer) {
		t.Errorf("the returned pointer is invalid")
		return
	}
}

func TestAdapter_withRemaining_Success(t *testing.T) {
	pointer := NewPointerForTests(0, 23)
	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retPointer, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining is invalid")
		return
	}

	if !reflect.DeepEqual(pointer, retPointer) {
		t.Errorf("the returned pointer is invalid")
		return
	}
}
