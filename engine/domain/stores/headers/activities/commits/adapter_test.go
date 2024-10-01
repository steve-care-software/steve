package commits

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

func TestAdapter_single_Success(t *testing.T) {
	commit := NewCommitForTests(
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
			modifications.NewModificationWithSaveForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
			modifications.NewModificationWithDeleteForTests("myDelete"),
		}),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(commit)
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

	if !reflect.DeepEqual(commit, retInstance) {
		t.Errorf("the returned commit is invalid")
		return
	}
}

func TestAdapter_single_withRemaining_Success(t *testing.T) {
	commit := NewCommitForTests(
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
			modifications.NewModificationWithSaveForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
			modifications.NewModificationWithDeleteForTests("myDelete"),
		}),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(commit)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retInstance, retRemaining, err := adapter.BytesToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retRemaining, remaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(commit, retInstance) {
		t.Errorf("the returned commit is invalid")
		return
	}
}

func TestAdapter_multiple_Success(t *testing.T) {
	commits := NewCommitsForTests([]Commit{
		NewCommitForTests(
			modifications.NewModificationsForTests([]modifications.Modification{
				modifications.NewModificationWithInsertForTests(
					resources.NewResourceForTests(
						"myIdentifier",
						pointers.NewPointerForTests(0, 23),
					),
				),
				modifications.NewModificationWithSaveForTests(
					resources.NewResourceForTests(
						"myIdentifier",
						pointers.NewPointerForTests(0, 23),
					),
				),
			}),
		),
		NewCommitForTests(
			modifications.NewModificationsForTests([]modifications.Modification{
				modifications.NewModificationWithDeleteForTests("myDelete"),
			}),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(commits)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(commits, retInstance) {
		t.Errorf("the returned commit is invalid")
		return
	}
}

func TestAdapter_multiple_withRemaining_Success(t *testing.T) {
	first := NewCommitForTests(
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithInsertForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
			modifications.NewModificationWithSaveForTests(
				resources.NewResourceForTests(
					"myIdentifier",
					pointers.NewPointerForTests(0, 23),
				),
			),
		}),
	)

	commits := NewCommitsForTests([]Commit{
		first,
		NewCommitWithParentForTests(
			modifications.NewModificationsForTests([]modifications.Modification{
				modifications.NewModificationWithDeleteForTests("myDelete"),
			}),
			first.Hash(),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(commits)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retInstance, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(retRemaining, remaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(commits, retInstance) {
		t.Errorf("the returned commit is invalid")
		return
	}
}
