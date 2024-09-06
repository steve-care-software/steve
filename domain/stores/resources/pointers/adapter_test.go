package pointers

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAdapter_single_Success(t *testing.T) {
	pointer := NewPointerForTests(0, 23)
	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPointer, retRemaining, err := adapter.BytesToInstance(retBytes)
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

func TestAdapter_single_withRemaining_Success(t *testing.T) {
	pointer := NewPointerForTests(0, 23)
	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(pointer)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retPointer, retRemaining, err := adapter.BytesToInstance(append(retBytes, remaining...))
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

func TestAdapter_multiple_Success(t *testing.T) {
	pointers := NewPointersForTests([]Pointer{
		NewPointerForTests(0, 23),
		NewPointerForTests(23, 2),
		NewPointerForTests(25, 29),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(pointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retPointers, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(pointers, retPointers) {
		t.Errorf("the returned pointers is invalid")
		return
	}
}

func TestAdapter_multiple_withRemaining_Success(t *testing.T) {
	pointers := NewPointersForTests([]Pointer{
		NewPointerForTests(0, 23),
		NewPointerForTests(23, 2),
		NewPointerForTests(25, 29),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(pointers)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retPointers, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining is invalid")
		return
	}

	if !reflect.DeepEqual(pointers, retPointers) {
		t.Errorf("the returned pointers is invalid")
		return
	}
}
