package roots

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/engine/domain/hash"
)

func TestAdapter_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pCommit, err := hashAdapter.FromBytes([]byte("owner"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	ownerPubKey, _, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	root := NewRootForTests(456, ownerPubKey, *pCommit)

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
