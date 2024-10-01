package resources

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/databases/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

func TestAdapter_single_Success(t *testing.T) {
	resource := NewResourceForTests(
		"myIdentifier",
		pointers.NewPointerForTests(0, 23),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(resource)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retResource, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the returned resource is invalid")
		return
	}
}

func TestAdapter_single_withRemaining_Success(t *testing.T) {
	resource := NewResourceForTests(
		"myIdentifier",
		pointers.NewPointerForTests(0, 23),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(resource)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retResource, retRemaining, err := adapter.BytesToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining is invalid")
		return
	}

	if !reflect.DeepEqual(resource, retResource) {
		t.Errorf("the returned resource is invalid")
		return
	}
}

func TestAdapter_multiple_Success(t *testing.T) {
	resources := NewResourcesForTests([]Resource{
		NewResourceForTests(
			"firstIDentifier",
			pointers.NewPointerForTests(0, 23),
		),
		NewResourceForTests(
			"secondIdentifier",
			pointers.NewPointerForTests(23, 55),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(resources)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retResources, retRemaining, err := adapter.BytesToInstances(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the remaining was expected to be empty")
		return
	}

	if !reflect.DeepEqual(resources, retResources) {
		t.Errorf("the returned resources is invalid")
		return
	}
}

func TestAdapter_multiple_withRemaining_Success(t *testing.T) {
	resources := NewResourcesForTests([]Resource{
		NewResourceForTests(
			"firstIDentifier",
			pointers.NewPointerForTests(0, 23),
		),
		NewResourceForTests(
			"secondIdentifier",
			pointers.NewPointerForTests(23, 55),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(resources)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retResources, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining is invalid")
		return
	}

	if !reflect.DeepEqual(resources, retResources) {
		t.Errorf("the returned resources is invalid")
		return
	}
}
