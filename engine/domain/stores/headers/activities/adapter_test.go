package activities

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/engine/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

func TestAdapter_single_Success(t *testing.T) {
	head := commits.NewCommitForTests(
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithDeleteForTests("myDelete"),
		}),
	)

	activity := NewActivityForTests(
		commits.NewCommitsForTests([]commits.Commit{
			commits.NewCommitForTests(
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
			head,
		}),
		head.Hash(),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(activity)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retInstance, retRemaining, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) != 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(activity, retInstance) {
		t.Errorf("the returned activity is invalid")
		return
	}
}
