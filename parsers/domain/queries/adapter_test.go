package queries

import (
	"bytes"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	adapter, err := NewAdapterFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	input := []byte(`
		v1;
		name: mySelector;
		myChain[0][0]->myChain[0][0]->myChain[0][0]->RULE;`)

	remaining := []byte("this is the remaining")

	retQuery, retRemaining, err := adapter.ToQuery(append(input, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if retQuery.Version() != 1 {
		t.Errorf("the version was expected to be %d, %d returned", 1, retQuery.Version())
		return
	}

	if retQuery.Name() != "mySelector" {
		t.Errorf("the name was expected to be '%s', '%s' returned", "mySelector", retQuery.Name())
		return
	}
}
