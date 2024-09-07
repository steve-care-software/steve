package modifications

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

func TestAdapter_list_Success(t *testing.T) {
	list := NewModificationsForTests([]Modification{
		NewModificationWithInsertForTests(
			resources.NewResourceForTests(
				"myIdentifier",
				pointers.NewPointerForTests(0, 23),
			),
		),
		NewModificationWithSaveForTests(
			resources.NewResourceForTests(
				"myIdentifier",
				pointers.NewPointerForTests(0, 23),
			),
		),
		NewModificationWithDeleteForTests("myDelete"),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(list)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstances, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(list, retInstances) {
		t.Errorf("the returned modification is invalid")
		return
	}
}

func TestAdapter_list_withRemaining_Success(t *testing.T) {
	list := NewModificationsForTests([]Modification{
		NewModificationWithInsertForTests(
			resources.NewResourceForTests(
				"myIdentifier",
				pointers.NewPointerForTests(0, 23),
			),
		),
		NewModificationWithSaveForTests(
			resources.NewResourceForTests(
				"myIdentifier",
				pointers.NewPointerForTests(0, 23),
			),
		),
		NewModificationWithDeleteForTests("myDelete"),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(list)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retInstances, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retRemaining, remaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(list, retInstances) {
		t.Errorf("the returned modification is invalid")
		return
	}
}

func TestAdapter_single_withInsert_Success(t *testing.T) {
	modification := NewModificationWithInsertForTests(
		resources.NewResourceForTests(
			"myIdentifier",
			pointers.NewPointerForTests(0, 23),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(modification)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(modification, retInstance) {
		t.Errorf("the returned modification is invalid")
		return
	}
}

func TestAdapter_single_withSave_Success(t *testing.T) {
	modification := NewModificationWithSaveForTests(
		resources.NewResourceForTests(
			"myIdentifier",
			pointers.NewPointerForTests(0, 23),
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(modification)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(modification, retInstance) {
		t.Errorf("the returned modification is invalid")
		return
	}
}

func TestAdapter_single_withDelete_Success(t *testing.T) {
	modification := NewModificationWithDeleteForTests("myDelete")
	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(modification)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(modification, retInstance) {
		t.Errorf("the returned modification is invalid")
		return
	}
}
