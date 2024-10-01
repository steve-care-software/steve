package headers

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

func TestAdapter_Success(t *testing.T) {
	header := NewHeaderForTests(
		resources.NewResourcesForTests([]resources.Resource{
			resources.NewResourceForTests(
				"firstIDentifier",
				pointers.NewPointerForTests(0, 23),
			),
			resources.NewResourceForTests(
				"secondIdentifier",
				pointers.NewPointerForTests(23, 55),
			),
		}),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(header)
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

	if !reflect.DeepEqual(header, retInstance) {
		t.Errorf("the returned header is invalid")
		return
	}
}

func TestAdapter_withActivity_Success(t *testing.T) {
	head := commits.NewCommitForTests(
		modifications.NewModificationsForTests([]modifications.Modification{
			modifications.NewModificationWithDeleteForTests("myDelete"),
		}),
	)

	header := NewHeaderWithActivityForTests(
		resources.NewResourcesForTests([]resources.Resource{
			resources.NewResourceForTests(
				"firstIDentifier",
				pointers.NewPointerForTests(0, 23),
			),
			resources.NewResourceForTests(
				"secondIdentifier",
				pointers.NewPointerForTests(23, 55),
			),
		}),
		activities.NewActivityForTests(
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
		),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(header)
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

	if !reflect.DeepEqual(header, retInstance) {
		t.Errorf("the returned header is invalid")
		return
	}
}
