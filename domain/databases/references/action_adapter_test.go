package references

import (
	"reflect"
	"testing"
)

func TestActionAdapter_withInsert_Success(t *testing.T) {
	action := NewActionWithInsert()
	adapter := NewActionAdapter()
	content, err := adapter.ToContent(action)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAction, err := adapter.ToAction(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(action, retAction) {
		t.Errorf("the returned action is invalid")
		return
	}
}

func TestActionAdapter_withDelete_Success(t *testing.T) {
	action := NewActionWithDelete()
	adapter := NewActionAdapter()
	content, err := adapter.ToContent(action)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAction, err := adapter.ToAction(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(action, retAction) {
		t.Errorf("the returned action is invalid")
		return
	}
}

func TestActionAdapter_withInsert_withDelete_Success(t *testing.T) {
	action := NewActionWithInsertAndDelete()
	adapter := NewActionAdapter()
	content, err := adapter.ToContent(action)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retAction, err := adapter.ToAction(content)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !reflect.DeepEqual(action, retAction) {
		t.Errorf("the returned action is invalid")
		return
	}
}
