package rules

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAdapter_withRemaining_Success(t *testing.T) {
	rules := NewRulesForTests(0, 2, 0.01)
	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(rules)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retRules, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(rules, retRules) {
		t.Errorf("the returned rules is invalid")
		return
	}
}
