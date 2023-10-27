package references

import (
	"reflect"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	reference := NewReferenceWithContentKeysForTests(0)
	adapter := NewAdapter()
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference, retReference) {
		t.Errorf("the returned reference is invalid")
		return
	}
}

func TestAdapter_withContentKeys_Success(t *testing.T) {
	reference := NewReferenceForTests(0)
	adapter := NewAdapter()
	content, err := adapter.ToContent(reference)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retReference, err := adapter.ToReference(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(reference, retReference) {
		t.Errorf("the returned reference is invalid")
		return
	}
}
