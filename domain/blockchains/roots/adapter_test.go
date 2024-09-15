package roots

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/domain/hash"
)

func TestAdapter_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pOwner, err := hashAdapter.FromBytes([]byte("owner"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pCommit, err := hashAdapter.FromBytes([]byte("owner"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	root := NewRootForTests(456, *pOwner, *pCommit)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(root)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retRoot, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(root, retRoot) {
		t.Errorf("the returned root is invalid")
		return
	}
}
